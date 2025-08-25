package api

import (
	"encoding/json"
	"fmt"
)

// Data defines an interface for all data objects from Synology API.
type Data any

// Response defines an interface for all responses from Synology API.
type Response interface {
	Data
}

// This is returned if no OTP code is provided when required: `{"error":{"code":403,"errors":{"token":"<some_token>","types":[{"type":"otp"}]}},"success":false}`
type ApiResponsePartialAuth[TData Data] struct {
	Success bool `json:"success"`
	Error   struct {
		Code   int32 `json:"code"`
		Errors struct {
			Token string `json:"token"`
			Types []struct {
				Type string `json:"type"`
			} `json:"types"`
		} `json:"errors"`
	} `json:"error"`
}

// ApiResponse is a concrete Response implementation.
// It is a generic struct with common to all Synology response fields.
type ApiResponse[TData Data] struct {
	Success bool     `json:"success"`
	Data    TData    `json:"data,omitempty"`
	Error   ApiError `json:"error"`
}

// UnmarshalJSON implements custom JSON unmarshalling for ApiResponse.
// This method handles the generic Data field and ensures proper error handling.
func (r *ApiResponse[TData]) UnmarshalJSON(data []byte) error {
	// Create a temporary struct to unmarshal the basic fields
	type Alias ApiResponse[TData]
	temp := &struct {
		Success bool            `json:"success"`
		Data    json.RawMessage `json:"data,omitempty"`
		Error   ApiError        `json:"error"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(data, temp); err != nil {
		return fmt.Errorf("failed to unmarshal ApiResponse: %w", err)
	}

	// Set the basic fields
	r.Success = temp.Success
	r.Error = temp.Error

	// Handle the Data field specially
	if len(temp.Data) > 0 && string(temp.Data) != "null" {
		// Create a new instance of TData
		var dataInstance TData
		if err := json.Unmarshal(temp.Data, &dataInstance); err != nil {
			return fmt.Errorf("failed to unmarshal data field: %w", err)
		}
		r.Data = dataInstance
	}

	return nil
}

// func NewApiResponse[TData Data]() *ApiResponse[TData] {
// 	return &ApiResponse[TData]{
// 		Data: new(TData),
// 	}
// }

// func NewApiResponseWithData[TData Data](data *TData) *ApiResponse[TData] {
// 	return &ApiResponse[TData]{
// 		Data: data,
// 	}
// }
