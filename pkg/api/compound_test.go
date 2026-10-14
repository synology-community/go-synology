package api

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/go-synology/pkg/query"
)

// TestCompoundElementMarshalJSON_Flat is the core encoding requirement from
// PLAT-705: a compound element's params merge directly onto the same JSON
// object as api/method/version. There must be no nested "params" wrapper --
// DSM's compound endpoint does not accept one.
func TestCompoundElementMarshalJSON_Flat(t *testing.T) {
	elem := CompoundElement{
		API:     "SYNO.Core.Share.Permission",
		Method:  "set",
		Version: 1,
		Params: map[string]any{
			"name":            "myshare",
			"user_group_type": "local_user",
		},
	}

	b, err := json.Marshal(elem)
	require.NoError(t, err)

	var obj map[string]any
	require.NoError(t, json.Unmarshal(b, &obj))

	require.Equal(t, "SYNO.Core.Share.Permission", obj["api"])
	require.Equal(t, "set", obj["method"])
	require.Equal(t, float64(1), obj["version"])
	require.Equal(t, "myshare", obj["name"])
	require.Equal(t, "local_user", obj["user_group_type"])

	// The defect this guards against: a nested "params" key instead of a
	// flat merge. DSM 403s (or ignores) that shape.
	_, hasParams := obj["params"]
	require.False(t, hasParams, "compound element must not nest params under a \"params\" key")

	// Exactly the five keys above -- nothing else leaked in or was dropped.
	require.Len(t, obj, 5)
}

// TestCompoundElementMarshalJSON_NilParams covers a sub-request that takes
// no parameters -- only api/method/version should appear.
func TestCompoundElementMarshalJSON_NilParams(t *testing.T) {
	elem := CompoundElement{API: "SYNO.Core.System", Method: "info", Version: 1}

	b, err := json.Marshal(elem)
	require.NoError(t, err)

	var obj map[string]any
	require.NoError(t, json.Unmarshal(b, &obj))
	require.Len(t, obj, 3)
	require.Equal(t, "SYNO.Core.System", obj["api"])
	require.Equal(t, "info", obj["method"])
	require.Equal(t, float64(1), obj["version"])
}

// TestCompoundElementMarshalJSON_NonObjectParams rejects a Params value that
// cannot flatten onto the element (a scalar/slice), rather than silently
// dropping it.
func TestCompoundElementMarshalJSON_NonObjectParams(t *testing.T) {
	elem := CompoundElement{API: "SYNO.Core.System", Method: "info", Version: 1, Params: []int{1, 2}}
	_, err := json.Marshal(elem)
	require.Error(t, err)
}

// TestCompoundRequestEncoding verifies the request-level form encoding: the
// "compound" value is a JSON array of the flat element shape, exactly the
// wire format verified against the live NAS in PLAT-705.
func TestCompoundRequestEncoding(t *testing.T) {
	req := CompoundRequest{
		Compound: []CompoundElement{
			{
				API:     "SYNO.Core.Share.Permission",
				Method:  "set",
				Version: 1,
				Params: map[string]any{
					"name":            "myshare",
					"user_group_type": "local_user",
					"permissions": []map[string]any{
						{
							"name":        "someuser",
							"is_readonly": false,
							"is_writable": true,
							"is_deny":     false,
							"is_custom":   false,
						},
					},
				},
			},
		},
		Mode:          "sequential",
		StopWhenError: false,
	}

	values, err := query.Values(req)
	require.NoError(t, err)

	require.Equal(t, "sequential", values.Get("mode"))
	require.Equal(t, "false", values.Get("stop_when_error"))

	var arr []map[string]any
	require.NoError(t, json.Unmarshal([]byte(values.Get("compound")), &arr),
		"compound must be a JSON array: got %s", values.Get("compound"))
	require.Len(t, arr, 1)

	sub := arr[0]
	require.Equal(t, "SYNO.Core.Share.Permission", sub["api"])
	require.Equal(t, "set", sub["method"])
	require.Equal(t, float64(1), sub["version"])
	require.Equal(t, "myshare", sub["name"])
	require.Equal(t, "local_user", sub["user_group_type"])
	_, hasParams := sub["params"]
	require.False(t, hasParams, "sub-request params must be flat, not nested under \"params\"")

	perms, ok := sub["permissions"].([]any)
	require.True(t, ok, "permissions must decode as an array: got %T", sub["permissions"])
	require.Len(t, perms, 1)
	row, ok := perms[0].(map[string]any)
	require.True(t, ok)
	require.Equal(t, false, row["is_readonly"])
	require.Equal(t, true, row["is_writable"])
	require.Equal(t, false, row["is_deny"])
	require.Equal(t, false, row["is_custom"])
}

func TestCompoundError(t *testing.T) {
	tests := []struct {
		name    string
		resp    *CompoundResponse
		wantErr bool
		check   func(t *testing.T, err error)
	}{
		{
			name:    "nil response",
			resp:    nil,
			wantErr: false,
		},
		{
			name: "all succeeded",
			resp: &CompoundResponse{
				HasFail: false,
				Result: []CompoundResult{
					{API: "SYNO.Core.Share.Permission", Method: "set", Success: true},
				},
			},
			wantErr: false,
		},
		{
			// The case this whole ticket exists for: envelope-level success
			// with a failed sub-request. Must error, not silently succeed.
			name: "envelope success, sub-request failed",
			resp: &CompoundResponse{
				HasFail: true,
				Result: []CompoundResult{
					{
						API:     "SYNO.Core.Share.Permission",
						Method:  "set",
						Success: false,
						Error:   ApiError{Code: 105},
					},
				},
			},
			wantErr: true,
			check: func(t *testing.T, err error) {
				var cfe CompoundFailedError
				require.ErrorAs(t, err, &cfe)
				require.Equal(t, 0, cfe.Index)
				require.Equal(t, "SYNO.Core.Share.Permission", cfe.API)
				require.Equal(t, "set", cfe.Method)
				require.Error(t, cfe.Unwrap())
			},
		},
		{
			name: "failed sub-request with no error code",
			resp: &CompoundResponse{
				HasFail: true,
				Result: []CompoundResult{
					{API: "SYNO.Core.Share.Permission", Method: "set", Success: false},
				},
			},
			wantErr: true,
		},
		{
			// Defensive: has_fail set but nothing in result[] says so. Must
			// still error rather than reading a partial/inconsistent
			// envelope as a clean success.
			name: "has_fail true with no failing result entries",
			resp: &CompoundResponse{
				HasFail: true,
				Result: []CompoundResult{
					{API: "SYNO.Core.Share.Permission", Method: "set", Success: true},
				},
			},
			wantErr: true,
		},
		{
			name:    "empty result, has_fail false",
			resp:    &CompoundResponse{HasFail: false, Result: nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CompoundError(tt.resp)
			if tt.wantErr {
				require.Error(t, err)
				if tt.check != nil {
					tt.check(t, err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCompoundFailedErrorUnwrap(t *testing.T) {
	inner := errors.New("boom")
	e := CompoundFailedError{Index: 2, API: "a", Method: "m", Err: inner}
	require.ErrorIs(t, e, inner)
	require.Contains(t, e.Error(), "boom")
	require.Contains(t, e.Error(), "a.m")
}
