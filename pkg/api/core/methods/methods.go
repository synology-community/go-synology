package methods

import (
	"github.com/synology-community/synology-api/pkg/api"
)

const (
	API_Core_System  = "SYNO.Core.System"
	API_Core_Package = "SYNO.Core.Package"
)

var (
	SystemInfo = api.Method{
		API:          API_Core_System,
		Version:      1,
		Method:       api.MethodInfo,
		ErrorSummary: api.GlobalErrors,
	}
	PackageList = api.Method{
		API:          API_Core_Package,
		Version:      2,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
)
