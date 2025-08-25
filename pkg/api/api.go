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
	Login(ctx context.Context, options LoginOptions) (*LoginResponse, error)
	GetApiInfo(ctx context.Context) (*ApiInfo, error)
	Password() string
}

const (
	Entry_Request   = "SYNO.Entry.Request"
	API_Auth        = "SYNO.API.Auth"
	API_Info        = "SYNO.API.Info"
	Core_NormalUser = "SYNO.Core.NormalUser"
)

var (
	Api_Info = Method{
		API:            API_Info,
		Version:        1,
		Method:         "query",
		ErrorSummaries: GlobalErrors,
	}
	Login = Method{
		API:            API_Auth,
		Version:        7,
		Method:         "login",
		ErrorSummaries: loginErrors,
	}
	Compound = Method{
		API:            Entry_Request,
		Version:        1,
		Method:         "request",
		ErrorSummaries: GlobalErrors,
	}
	Core_UserInfo = Method{
		API:            Core_NormalUser,
		Version:        1,
		Method:         "get",
		ErrorSummaries: GlobalErrors,
	}
)

var API_METHODS = APIMethodLookup{
	"Login": {
		API:            API_Auth,
		Version:        7,
		Method:         "login",
		ErrorSummaries: GlobalErrors,
	},
}
