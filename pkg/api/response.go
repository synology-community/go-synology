package api

// Data defines an interface for all data objects from Synology API.
type Data any

// Response defines an interface for all responses from Synology API.
type Response interface {
	Data
}

// ApiResponse is a concrete Response implementation.
// It is a generic struct with common to all Synology response fields.
type ApiResponse[TData Data] struct {
	Success bool     `json:"success"`
	Data    TData    `json:"data,omitempty"`
	Error   ApiError `json:"error"`
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
