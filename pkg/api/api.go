package api

import "context"

type API interface {
	Login(ctx context.Context, user, password string) (*LoginResponse, error)
}

var API_METHODS = APIMethodLookup{
	"Login": {
		API:          "SYNO.API.Auth",
		Version:      7,
		Method:       "login",
		ErrorSummary: GlobalErrors,
	},
}
