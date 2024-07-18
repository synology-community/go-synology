package methods

import "github.com/synology-community/go-synology/pkg/api"

const (
	API_Portal  = "SYNO.WebStation.WebService.Portal"
	API_Service = "SYNO.WebStation.WebService.Service"
)

var (
	PortalCreate = api.Method{
		API:            API_Portal,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}

	PortalUpdate = api.Method{
		API:            API_Portal,
		Version:        1,
		Method:         api.MethodUpdate,
		ErrorSummaries: api.GlobalErrors,
	}

	PortalList = api.Method{
		API:            API_Portal,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}

	ServiceList = api.Method{
		API:            API_Service,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
)
