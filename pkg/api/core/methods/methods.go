package methods

import (
	"github.com/synology-community/synology-api/pkg/api"
)

const (
	API_Core_System               = "SYNO.Core.System"
	API_Core_Package              = "SYNO.Core.Package"
	API_Core_Package_Installation = "SYNO.Core.Package.Installation"
	API_Core_Package_Server       = "SYNO.Core.Package.Server"
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
	PackageGet = api.Method{
		API:          API_Core_Package,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
	PackageServerList = api.Method{
		API:          API_Core_Package_Server,
		Version:      2,
		Method:       api.MethodList,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationInstall = api.Method{
		API:          API_Core_Package_Installation,
		Version:      1,
		Method:       api.MethodInstall,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationStatus = api.Method{
		API:          API_Core_Package_Installation,
		Version:      1,
		Method:       api.MethodStatus,
		ErrorSummary: api.GlobalErrors,
	}
)
