package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApiError_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name              string
		jsonData          string
		expectedCode      int
		expectedErrors    []ErrorFields
		shouldHaveSummary bool
		wantError         bool
	}{
		{
			name:         "Valid error with known code and single error field",
			jsonData:     `{"code": 101, "errors": [{"code": 407, "path": "test/example"}]}`,
			expectedCode: 101,
			expectedErrors: []ErrorFields{
				{Code: 407, Fields: map[string]any{"path": "test/example"}},
			},
			shouldHaveSummary: true,
			wantError:         false,
		},
		{
			name:         "Valid error with multiple error fields",
			jsonData:     `{"code": 102, "errors": [{"code": 407, "field1": "error1"}, {"code": 408, "field2": "error2"}]}`,
			expectedCode: 102,
			expectedErrors: []ErrorFields{
				{Code: 407, Fields: map[string]any{"field1": "error1"}},
				{Code: 408, Fields: map[string]any{"field2": "error2"}},
			},
			shouldHaveSummary: true,
			wantError:         false,
		},
		{
			name:              "Valid error with unknown code",
			jsonData:          `{"code": 9999}`,
			expectedCode:      9999,
			expectedErrors:    nil,
			shouldHaveSummary: true, // Should still have summary, even if "Unknown error code"
			wantError:         false,
		},
		{
			name:              "Error with zero code",
			jsonData:          `{"code": 0}`,
			expectedCode:      0,
			expectedErrors:    nil,
			shouldHaveSummary: false, // Zero code should not populate summary
			wantError:         false,
		},
		{
			name:         "Error with complex errors field",
			jsonData:     `{"code": 102, "errors": [{"code": 409, "field1": {"nested": "value"}}, {"code": 410, "field2": 123}]}`,
			expectedCode: 102,
			expectedErrors: []ErrorFields{
				{Code: 409, Fields: map[string]any{"field1": map[string]any{"nested": "value"}}},
				{Code: 410, Fields: map[string]any{"field2": float64(123)}},
			},
			shouldHaveSummary: true,
			wantError:         false,
		},
		{
			name:              "Empty JSON object",
			jsonData:          `{}`,
			expectedCode:      0,
			expectedErrors:    nil,
			shouldHaveSummary: false,
			wantError:         false,
		},
		{
			name:              "Only code field",
			jsonData:          `{"code": 103}`,
			expectedCode:      103,
			expectedErrors:    nil,
			shouldHaveSummary: true,
			wantError:         false,
		},
		{
			name:              "Empty errors array",
			jsonData:          `{"code": 104, "errors": []}`,
			expectedCode:      104,
			expectedErrors:    []ErrorFields{},
			shouldHaveSummary: true,
			wantError:         false,
		},
		{
			name:         "Real API error response with code 1100",
			jsonData:     `{"code": 1100, "errors": [{"code": 407, "path": "test/foobar"}]}`,
			expectedCode: 1100,
			expectedErrors: []ErrorFields{
				{Code: 407, Fields: map[string]any{"path": "test/foobar"}},
			},
			shouldHaveSummary: true,
			wantError:         false,
		},
		{
			name:      "Invalid JSON",
			jsonData:  `{"code": invalid}`,
			wantError: true,
		},
		{
			name:      "Invalid code type",
			jsonData:  `{"code": "not_a_number"}`,
			wantError: true,
		},
		{
			name:      "Invalid errors structure (missing code field)",
			jsonData:  `{"code": 101, "errors": [{"field1": "error1"}]}`,
			wantError: true,
		},
		{
			name:      "Invalid errors structure (not array)",
			jsonData:  `{"code": 101, "errors": {"code": 407, "field1": "error1"}}`,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var apiError ApiError
			err := json.Unmarshal([]byte(tt.jsonData), &apiError)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedCode, apiError.Code)

			if tt.expectedErrors != nil {
				assert.Equal(t, tt.expectedErrors, apiError.Errors)
			} else {
				assert.Empty(t, apiError.Errors)
			}

			if tt.shouldHaveSummary {
				assert.NotEmpty(
					t,
					apiError.Summary,
					"Summary should be populated for non-zero error codes",
				)
				if tt.expectedCode == 101 {
					assert.Equal(t, "No parameter of API, method or version", apiError.Summary)
				} else if tt.expectedCode == 102 {
					assert.Equal(t, "The requested API does not exist", apiError.Summary)
				} else if tt.expectedCode == 103 {
					assert.Equal(t, "The requested method does not exist", apiError.Summary)
				} else if tt.expectedCode == 9999 {
					assert.Equal(t, "Unknown error code", apiError.Summary)
				}
			} else {
				assert.Empty(t, apiError.Summary, "Summary should not be populated for zero error code")
			}
		})
	}
}

func TestApiError_UnmarshalJSON_WithRealApiResponse(t *testing.T) {
	// Test unmarshalling ApiError as part of a full API response
	jsonResponse := `{
		"success": false,
		"error": {
			"code": 105,
			"errors": [
				{"code": 407, "permission": "denied"},
				{"code": 408, "session": "invalid"}
			]
		}
	}`

	var response ApiResponse[any]
	err := json.Unmarshal([]byte(jsonResponse), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, 105, response.Error.Code)
	assert.Equal(t, "The logged in session does not have permission", response.Error.Summary)
	assert.Equal(t, []ErrorFields{
		{Code: 407, Fields: map[string]any{"permission": "denied"}},
		{Code: 408, Fields: map[string]any{"session": "invalid"}},
	}, response.Error.Errors)
}

func TestApiError_UnmarshalJSON_WithCode1100Response(t *testing.T) {
	// Test the specific case provided by the user
	jsonResponse := `{
		"error": {
			"code": 1100,
			"errors": [
				{
					"code": 407,
					"path": "test/foobar"
				}
			]
		},
		"success": false
	}`

	var response ApiResponse[any]
	err := json.Unmarshal([]byte(jsonResponse), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, 1100, response.Error.Code)
	assert.Equal(t, "Unknown error code", response.Error.Summary) // 1100 is not in global errors

	require.Len(t, response.Error.Errors, 1)
	errorField := response.Error.Errors[0]
	assert.Equal(t, 407, errorField.Code)
	assert.Equal(t, "test/foobar", errorField.Fields["path"])
}

func TestApiError_UnmarshalJSON_PreservesOtherFields(t *testing.T) {
	// Test that unmarshalling doesn't interfere with other fields
	apiError := ApiError{
		Code:    101,
		Summary: "Original summary", // This should be overwritten
		Errors:  []ErrorFields{{Code: 406, Fields: map[string]any{"old": "value"}}},
	}

	jsonData := `{"code": 102, "errors": [{"code": 407, "new": "value"}]}`
	err := json.Unmarshal([]byte(jsonData), &apiError)
	require.NoError(t, err)

	assert.Equal(t, 102, apiError.Code)
	assert.Equal(
		t,
		"The requested API does not exist",
		apiError.Summary,
	) // Should be updated
	assert.Equal(
		t,
		[]ErrorFields{{Code: 407, Fields: map[string]any{"new": "value"}}},
		apiError.Errors,
	) // Should be updated
}

func TestApiError_UnmarshalJSON_ErrorHandling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
	}{
		{
			name:     "Invalid JSON structure",
			jsonData: `["code", 101]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var apiError ApiError
			err := json.Unmarshal([]byte(tt.jsonData), &apiError)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "failed to unmarshal ApiError")
		})
	}
}
