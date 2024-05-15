package client

import (
	"context"
	"net/url"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/util"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type Nil struct{}

func newClient(t *testing.T) *APIClient {
	c, err := New(Options{
		Host:       os.Getenv("SYNOLOGY_HOST"),
		VerifyCert: false,
	})
	if err != nil {
		t.Error(err)
		require.NoError(t, err)
	}

	if r, err := c.Login(context.Background(), os.Getenv("SYNOLOGY_USER"), os.Getenv("SYNOLOGY_PASSWORD")); err != nil {
		t.Error(err)
		require.NoError(t, err)
	} else {
		log.Infoln("Login successful")
		log.Infof("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)
	}

	if client, ok := c.(*APIClient); !ok {
		t.Error("Client is not of type APIClient")
		require.True(t, ok)
		return nil
	} else {
		return client
	}
}

func Test_FileStationClient_Upload(t *testing.T) {
	c := newClient(t)

	file := form.File{
		Name:    "test.txt",
		Content: "Hello, World!",
	}

	_, err := c.FileStationAPI().Upload(context.Background(), "/data/foo", file, true, true)
	require.NoError(t, err)
}

func Test_FileStationClient_MD5(t *testing.T) {
	c := newClient(t)

	file := form.File{
		Name:    "test.txt",
		Content: "Hello, World!",
	}

	_, err := c.FileStationAPI().Upload(context.Background(), "/data/foo", file, true, true)
	require.NoError(t, err)

	_, err = c.FileStationAPI().MD5(context.Background(), "/data/foo/test.txt")
	require.NoError(t, err)
}

func Test_FileStationClient_DeleteStart(t *testing.T) {
	c := newClient(t)

	_, err := c.FileStationAPI().DeleteStart(context.Background(), []string{"/data/foodbar"}, true)
	require.NoError(t, err)
}

func Test_FileStationClient_DeleteStatus(t *testing.T) {
	c := newClient(t)

	r, err := c.FileStationAPI().DeleteStart(context.Background(), []string{"/data/foodbar"}, true)
	require.NoError(t, err)

	_, err = c.FileStationAPI().DeleteStatus(context.Background(), r.TaskID)
	require.NoError(t, err)
}

func TestMarshalURL(t *testing.T) {
	type embeddedStruct struct {
		EmbeddedString string `form:"embedded_string" url:"embedded_string"`
		EmbeddedInt    int    `form:"embedded_int" url:"embedded_int"`
	}

	testCases := []struct {
		name     string
		in       any
		expected url.Values
	}{
		{
			name: "scalar types",
			in: struct {
				Name    string `form:"name" url:"name"`
				ID      int    `form:"id" url:"id"`
				Enabled bool   `form:"enabled" url:"enabled"`
			}{
				Name:    "name value",
				ID:      2,
				Enabled: true,
			},
			expected: url.Values{
				"name":    []string{"name value"},
				"id":      []string{"2"},
				"enabled": []string{"true"},
			},
		},
		{
			name: "slice types",
			in: struct {
				Names []string `form:"names" url:"names"`
				IDs   []int    `form:"ids" url:"ids"`
			}{
				Names: []string{"value 1", "value 2"},
				IDs:   []int{1, 2, 3},
			},
			expected: url.Values{
				"names": []string{"[\"value 1\",\"value 2\"]"},
				"ids":   []string{"[1,2,3]"},
			},
		},
		{
			name: "embedded struct",
			in: struct {
				embeddedStruct
				Name string `form:"name" url:"name"`
			}{
				embeddedStruct: embeddedStruct{
					EmbeddedString: "my string",
					EmbeddedInt:    5,
				},
				Name: "field name",
			},
			expected: url.Values{
				"name":            []string{"field name"},
				"embedded_string": []string{"my string"},
				"embedded_int":    []string{"5"},
			},
		},
		{
			name: "unexported field without tag",
			in: struct {
				Name       string `form:"name" url:"name"`
				ID         int    `form:"id" url:"id"`
				unexported string
			}{
				Name:       "name value",
				ID:         2,
				unexported: "must be skipped",
			},
			expected: url.Values{
				"name": []string{"name value"},
				"id":   []string{"2"},
			},
		},
		{
			name: "unexported field with tag",
			in: struct {
				Name       string `form:"name" url:"name"`
				ID         int    `form:"id" url:"id"`
				unexported string `form:"unexported" url:"unexported"`
			}{
				Name:       "name value",
				ID:         2,
				unexported: "with explicit tag",
			},
			expected: url.Values{
				"name":       []string{"name value"},
				"id":         []string{"2"},
				"unexported": []string{"with explicit tag"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := util.MarshalURL(tc.in)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestHandleErrors(t *testing.T) {
	globalErrors := api.ErrorSummary{
		100: "global error 100",
		101: "global error 101",
		102: "global error 102",
	}

	testCases := []struct {
		name                string
		response            api.ApiResponse[Nil]
		responseKnownErrors []api.ErrorSummary
		expected            api.ApiError
	}{
		{
			name: "global errors only",
			response: api.ApiResponse[Nil]{
				Success: false,
				Data:    Nil{},
				Error: api.ApiError{
					Code: 100,
					Errors: []api.ErrorItem{
						{Code: 101},
						{Code: 102, Details: api.ErrorFields{"path": "/some/path", "code": 102, "reason": "a reason"}},
					},
				},
			},
			expected: api.ApiError{
				Code:    100,
				Summary: "global error 100",
				Errors: []api.ErrorItem{
					{
						Code:    101,
						Summary: "global error 101",
					},
					{
						Code:    102,
						Summary: "global error 102",
						Details: api.ErrorFields{"path": "/some/path", "code": 102, "reason": "a reason"},
					},
				},
			},
		},
		{
			name: "response-specific error",
			response: api.ApiResponse[Nil]{
				Success: false,
				Data:    Nil{},
				Error: api.ApiError{
					Code: 100,
					Errors: []api.ErrorItem{
						{Code: 101},
						{Code: 202, Details: api.ErrorFields{"code": 202}},
					},
				},
			},
			responseKnownErrors: []api.ErrorSummary{
				{
					202: "response error 202",
				},
			},
			expected: api.ApiError{
				Code:    100,
				Summary: "global error 100",
				Errors: []api.ErrorItem{
					{
						Code:    101,
						Summary: "global error 101",
					},
					{
						Code:    202,
						Summary: "response error 202",
						Details: api.ErrorFields{"code": 202},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := handleErrors(tc.response,
				globalErrors,
			)
			assert.Error(t, actual)
		})
	}
}
