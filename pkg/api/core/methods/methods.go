package methods

import (
	"github.com/synology-community/go-synology/pkg/api"
)

const (
	Core_System                 = "SYNO.Core.System"
	Core_Package                = "SYNO.Core.Package"
	Core_Package_Feed           = "SYNO.Core.Package.Feed"
	Core_Package_Installation   = "SYNO.Core.Package.Installation"
	Core_Package_Uninstallation = "SYNO.Core.Package.Uninstallation"
	Core_Package_Server         = "SYNO.Core.Package.Server"
	Core_Package_Setting        = "SYNO.Core.Package.Setting"
	DSM_PortEnable              = "SYNO.DSM.PortEnable"
	Core_Share                  = "SYNO.Core.Share"
	Core_Storage_Volume         = "SYNO.Core.Storage.Volume"
)

var (
	SystemInfo = api.Method{
		API:          Core_System,
		Version:      1,
		Method:       api.MethodInfo,
		ErrorSummary: api.GlobalErrors,
	}
	PackageList = api.Method{
		API:          Core_Package,
		Version:      2,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
	PackageGet = api.Method{
		API:          Core_Package,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
	PackageServerList = api.Method{
		API:          Core_Package_Server,
		Version:      2,
		Method:       api.MethodList,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationInstall = api.Method{
		API:          Core_Package_Installation,
		Version:      1,
		Method:       api.MethodInstall,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationDelete = api.Method{
		API:          Core_Package_Installation,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: api.GlobalErrors,
	}
	PackageUnistallationUninstall = api.Method{
		API:          Core_Package_Uninstallation,
		Version:      1,
		Method:       api.MethodUninstall,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationStatus = api.Method{
		API:          Core_Package_Installation,
		Version:      1,
		Method:       api.MethodStatus,
		ErrorSummary: api.GlobalErrors,
	}
	PackageInstallationCheck = api.Method{
		API:          Core_Package_Installation,
		Version:      2,
		Method:       api.MethodCheck,
		ErrorSummary: api.GlobalErrors,
	}
	IsPortBlock = api.Method{
		API:          DSM_PortEnable,
		Version:      1,
		Method:       api.MethodIsPortBlock,
		ErrorSummary: api.GlobalErrors,
	}
	IsPkgEnable = api.Method{
		API:          DSM_PortEnable,
		Version:      1,
		Method:       "is_pkg_enable",
		ErrorSummary: api.GlobalErrors,
	}
	PackageFeedList = api.Method{
		API:          Core_Package_Feed,
		Version:      1,
		Method:       api.MethodList,
		ErrorSummary: api.GlobalErrors,
	}
	PackageFeedAdd = api.Method{
		API:          Core_Package_Feed,
		Version:      1,
		Method:       api.MethodAdd,
		ErrorSummary: api.GlobalErrors,
	}
	PackageFeedDelete = api.Method{
		API:          Core_Package_Feed,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: api.GlobalErrors,
	}
	PackageSettingGet = api.Method{
		API:          Core_Package_Setting,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
	ShareList = api.Method{
		API:          Core_Share,
		Version:      1,
		Method:       api.MethodList,
		ErrorSummary: api.GlobalErrors,
	}
	ShareGet = api.Method{
		API:          Core_Share,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
	ShareCreate = api.Method{
		API:          Core_Share,
		Version:      1,
		Method:       api.MethodCreate,
		ErrorSummary: api.GlobalErrors,
	}
	ShareDelete = api.Method{
		API:          Core_Share,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: api.GlobalErrors,
	}
	VolumeList = api.Method{
		API:          Core_Storage_Volume,
		Version:      1,
		Method:       api.MethodList,
		ErrorSummary: api.GlobalErrors,
	}
)
