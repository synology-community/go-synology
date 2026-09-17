package core

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/go-synology/pkg/api"
)

// newTestClient spins up a fake DSM that authenticates any login and then
// dispatches every other call to handleEntry. r.Form (populated by
// r.ParseForm) carries the request's parameters regardless of whether the
// underlying call was a GET with a query string (api.Get, used by
// SharePermissionList) or a form-encoded POST body (api.Post/postEntry,
// used for the SYNO.Entry.Request compound envelope SharePermissionSet
// needs) -- both land in r.Form uniformly, so handleEntry does not need to
// know which one it's looking at.
func newTestClient(
	t *testing.T,
	handleEntry func(w http.ResponseWriter, r *http.Request),
) Api {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, r.ParseForm())

		if r.Form.Get("api") == "SYNO.API.Auth" {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"data": map[string]any{
					"sid":       "test-sid",
					"synotoken": "test-token",
				},
			})
			return
		}

		handleEntry(w, r)
	}))
	t.Cleanup(srv.Close)

	c, err := api.New(api.Options{Host: srv.URL, AllowHTTP: true, VerifyCert: false})
	require.NoError(t, err)
	_, err = c.Login(context.Background(), api.LoginOptions{Username: "admin", Password: "pw"})
	require.NoError(t, err)

	return New(c)
}

// TestSharePermissionSet_EncodesFlatCompoundEnvelope is the encoding
// requirement from PLAT-705's verified wire format: a form-encoded POST to
// entry.cgi with api=SYNO.Entry.Request, method=request, version=1,
// mode=sequential, stop_when_error=false, and a "compound" value containing
// exactly one FLAT sub-request object -- no nested "params" wrapper, and
// _sid present only on the envelope, never inside the sub-request.
func TestSharePermissionSet_EncodesFlatCompoundEnvelope(t *testing.T) {
	var seenMethod string
	var seenForm map[string][]string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		seenMethod = r.Method
		seenForm = r.Form
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"data": map[string]any{
				"has_fail": false,
				"result": []map[string]any{
					{
						"api":     "SYNO.Core.Share.Permission",
						"method":  "set",
						"version": 1,
						"success": true,
						"data":    map[string]any{},
					},
				},
			},
		})
	})

	err := client.SharePermissionSet(context.Background(), "myshare", ShareUserGroupTypeLocalUser,
		[]SharePermissionSetEntry{
			{Name: "someuser", IsReadonly: false, IsWritable: true, IsDeny: false, IsCustom: false},
		})
	require.NoError(t, err)

	require.Equal(t, http.MethodPost, seenMethod, "the compound envelope must be POSTed, per PLAT-705")

	get := func(k string) string {
		v := seenForm[k]
		if len(v) == 0 {
			return ""
		}
		return v[0]
	}
	require.Equal(t, "SYNO.Entry.Request", get("api"))
	require.Equal(t, "request", get("method"))
	require.Equal(t, "1", get("version"))
	require.Equal(t, "sequential", get("mode"))
	require.Equal(t, "false", get("stop_when_error"))
	require.Equal(t, "test-sid", get("_sid"), "_sid must be present on the envelope")

	var arr []map[string]any
	require.NoError(t, json.Unmarshal([]byte(get("compound")), &arr))
	require.Len(t, arr, 1)

	sub := arr[0]
	require.Equal(t, "SYNO.Core.Share.Permission", sub["api"])
	require.Equal(t, "set", sub["method"])
	require.Equal(t, float64(1), sub["version"])
	require.Equal(t, "myshare", sub["name"])
	require.Equal(t, "local_user", sub["user_group_type"])

	_, hasParams := sub["params"]
	require.False(t, hasParams, "sub-request must be flat, not nested under a \"params\" key")
	_, hasSid := sub["_sid"]
	require.False(t, hasSid, "_sid must not appear inside the sub-request")

	perms, ok := sub["permissions"].([]any)
	require.True(t, ok)
	require.Len(t, perms, 1)
	row := perms[0].(map[string]any)
	require.Equal(t, "someuser", row["name"])
	require.Equal(t, false, row["is_readonly"])
	require.Equal(t, true, row["is_writable"])
	require.Equal(t, false, row["is_deny"])
	require.Equal(t, false, row["is_custom"])
}

// TestSharePermissionSet_PartialFailureErrors is the single most important
// case in PLAT-705: the compound envelope can report success:true at the
// top level while the wrapped SYNO.Core.Share.Permission set sub-request
// failed. SharePermissionSet must surface that as an error, not as success.
func TestSharePermissionSet_PartialFailureErrors(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true, // envelope-level success
			"data": map[string]any{
				"has_fail": true,
				"result": []map[string]any{
					{
						"api":     "SYNO.Core.Share.Permission",
						"method":  "set",
						"version": 1,
						"success": false,
						"error":   map[string]any{"code": 105},
					},
				},
			},
		})
	})

	err := client.SharePermissionSet(context.Background(), "myshare", ShareUserGroupTypeLocalUser,
		[]SharePermissionSetEntry{
			{Name: "someuser", IsReadonly: false, IsWritable: true, IsDeny: false, IsCustom: false},
		})

	require.Error(t, err, "envelope success:true with a failed sub-request must not read as success")

	var cfe api.CompoundFailedError
	require.ErrorAs(t, err, &cfe)
	require.Equal(t, "SYNO.Core.Share.Permission", cfe.API)
	require.Equal(t, "set", cfe.Method)
}

// TestSharePermissionSet_EnvelopeLevelFailure covers the ordinary failure
// path (the envelope itself reports success:false), which Post/handle
// already turns into an error before SharePermissionSet ever sees a
// CompoundResponse to check.
func TestSharePermissionSet_EnvelopeLevelFailure(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"error":   map[string]any{"code": 119},
		})
	})

	err := client.SharePermissionSet(context.Background(), "myshare", ShareUserGroupTypeLocalUser,
		[]SharePermissionSetEntry{{Name: "someuser", IsWritable: true}})
	require.Error(t, err)
}

// TestSharePermissionList_DirectCall covers the list side, which -- unlike
// set -- DSM accepts called directly as a plain GET, with no compound
// envelope. This matches the house pattern for ShareList/ShareGet
// (api.Get), not the POST that SharePermissionSet needs.
func TestSharePermissionList_DirectCall(t *testing.T) {
	var seenMethod, seenAPI, seenDSMMethod string
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		seenMethod = r.Method
		seenAPI = r.Form.Get("api")
		seenDSMMethod = r.Form.Get("method")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"data": map[string]any{
				"items": []map[string]any{
					{
						"name":        "someuser",
						"is_admin":    false,
						"is_custom":   false,
						"is_deny":     false,
						"is_readonly": false,
						"is_writable": true,
					},
				},
				"total": 1,
			},
		})
	})

	resp, err := client.SharePermissionList(context.Background(), "myshare", ShareUserGroupTypeLocalUser)
	require.NoError(t, err)

	require.Equal(t, http.MethodGet, seenMethod, "list is a direct GET, not a compound POST")
	require.Equal(t, "SYNO.Core.Share.Permission", seenAPI)
	require.Equal(t, "list", seenDSMMethod)

	require.Len(t, resp.Items, 1)
	require.Equal(t, "someuser", resp.Items[0].Name)
	require.True(t, resp.Items[0].IsWritable)
}

// A compound envelope can report success:true while carrying no results at
// all -- meaning DSM executed none of the sub-requests that were sent. Without
// a length check the caller's `if err != nil` reads that as "everything
// worked". This is the silent-pass shape behind PLAT-698, PLAT-711 and
// PLAT-715, and shipping it inside the fix for those would be perverse.
func TestSharePermissionSet_EmptyResultIsNotSuccess(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true, // envelope claims success
			"data": map[string]any{
				"has_fail": false,              // nothing flagged as failed
				"result":   []map[string]any{}, // ...because nothing ran
			},
		})
	})

	err := client.SharePermissionSet(context.Background(), "myshare", ShareUserGroupTypeLocalUser,
		[]SharePermissionSetEntry{
			{Name: "someuser", IsReadonly: false, IsWritable: true, IsDeny: false, IsCustom: false},
		})

	require.Error(t, err,
		"a compound that executed zero of one sub-request must not report success")
	require.Contains(t, err.Error(), "did not execute what was sent")
}
