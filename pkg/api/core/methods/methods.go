package methods

import (
	"github.com/synology-community/go-synology/pkg/api"
)

const (
	API_Core_System                 = "SYNO.Core.System"
	API_Core_Package                = "SYNO.Core.Package"
	API_Core_Package_Feed           = "SYNO.Core.Package.Feed"
	API_Core_Package_Installation   = "SYNO.Core.Package.Installation"
	API_Core_Package_Uninstallation = "SYNO.Core.Package.Uninstallation"
	API_Core_Package_Server         = "SYNO.Core.Package.Server"
	API_Core_Package_Setting        = "SYNO.Core.Package.Setting"
	API_DSM_PortEnable              = "SYNO.DSM.PortEnable"
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
	PackageInstallationDelete = api.Method{
		API:          API_Core_Package_Installation,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: api.GlobalErrors,
	}
	PackageUnistallationUninstall = api.Method{
		API:          API_Core_Package_Uninstallation,
		Version:      1,
		Method:       api.MethodUninstall,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationStatus = api.Method{
		API:          API_Core_Package_Installation,
		Version:      1,
		Method:       api.MethodStatus,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationCheck = api.Method{
		API:          API_Core_Package_Installation,
		Version:      2,
		Method:       api.MethodCheck,
		ErrorSummary: api.GlobalErrors,
	}
	IsPortBlock = api.Method{
		API:          API_DSM_PortEnable,
		Version:      1,
		Method:       api.MethodIsPortBlock,
		ErrorSummary: api.GlobalErrors,
	}
	IsPkgEnable = api.Method{
		API:          API_DSM_PortEnable,
		Version:      1,
		Method:       "is_pkg_enable",
		ErrorSummary: api.GlobalErrors,
	}
	PackageFeedList = api.Method{
		API:          API_Core_Package_Feed,
		Version:      1,
		Method:       api.MethodList,
		ErrorSummary: api.GlobalErrors,
	}
	PackageFeedAdd = api.Method{
		API:          API_Core_Package_Feed,
		Version:      1,
		Method:       api.MethodAdd,
		ErrorSummary: api.GlobalErrors,
	}
	PackageFeedDelete = api.Method{
		API:          API_Core_Package_Feed,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: api.GlobalErrors,
	}
	PackageSettingGet = api.Method{
		API:          API_Core_Package_Setting,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
)
