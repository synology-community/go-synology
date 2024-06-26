package api

import (
	"context"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

type Api interface {
	Client() *retryablehttp.Client
	BaseUrl() *url.URL
	Credentials() Credentials
	Login(ctx context.Context, user, password string) (*LoginResponse, error)
}

const (
	Entry_Request = "SYNO.Entry.Request"
	API_Auth      = "SYNO.API.Auth"
)

var (
	Login = Method{
		API:          API_Auth,
		Version:      7,
		Method:       "login",
		ErrorSummary: GlobalErrors,
	}
	Compound = Method{
		API:          Entry_Request,
		Version:      1,
		Method:       "request",
		ErrorSummary: GlobalErrors,
	}
)

var API_METHODS = APIMethodLookup{
	"Login": {
		API:          API_Auth,
		Version:      7,
		Method:       "login",
		ErrorSummary: GlobalErrors,
	},
}
