package api

import (
	"context"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

type Api interface {
	Client() *retryablehttp.Client
	BaseUrl() url.URL
	Credentials() Credentials
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
