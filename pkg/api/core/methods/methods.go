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
	Core_TaskScheduler          = "SYNO.Core.TaskScheduler"
	Core_Task_Root              = "SYNO.Core.TaskScheduler.Root"
	Core_Password_Confirm       = "SYNO.Core.User.PasswordConfirm"
)

var (
	SystemInfo = api.Method{
		API:            Core_System,
		Version:        1,
		Method:         api.MethodInfo,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageList = api.Method{
		API:            Core_Package,
		Version:        2,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageGet = api.Method{
		API:            Core_Package,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageServerList = api.Method{
		API:            Core_Package_Server,
		Version:        2,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageInstallationInstall = api.Method{
		API:            Core_Package_Installation,
		Version:        1,
		Method:         api.MethodInstall,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageInstallationDelete = api.Method{
		API:            Core_Package_Installation,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageUnistallationUninstall = api.Method{
		API:            Core_Package_Uninstallation,
		Version:        1,
		Method:         api.MethodUninstall,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageInstallationStatus = api.Method{
		API:            Core_Package_Installation,
		Version:        1,
		Method:         api.MethodStatus,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageInstallationCheck = api.Method{
		API:            Core_Package_Installation,
		Version:        2,
		Method:         api.MethodCheck,
		ErrorSummaries: api.GlobalErrors,
	}
	IsPortBlock = api.Method{
		API:            DSM_PortEnable,
		Version:        1,
		Method:         api.MethodIsPortBlock,
		ErrorSummaries: api.GlobalErrors,
	}
	IsPkgEnable = api.Method{
		API:            DSM_PortEnable,
		Version:        1,
		Method:         "is_pkg_enable",
		ErrorSummaries: api.GlobalErrors,
	}
	PackageFeedList = api.Method{
		API:            Core_Package_Feed,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageFeedAdd = api.Method{
		API:            Core_Package_Feed,
		Version:        1,
		Method:         api.MethodAdd,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageFeedDelete = api.Method{
		API:            Core_Package_Feed,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	PackageSettingGet = api.Method{
		API:            Core_Package_Setting,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
	ShareList = api.Method{
		API:            Core_Share,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	ShareGet = api.Method{
		API:            Core_Share,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
	ShareCreate = api.Method{
		API:            Core_Share,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	ShareDelete = api.Method{
		API:            Core_Share,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	VolumeList = api.Method{
		API:            Core_Storage_Volume,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	TaskList = api.Method{
		API:            Core_TaskScheduler,
		Version:        3,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	TaskGet = api.Method{
		API:            Core_TaskScheduler,
		Version:        3,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
	TaskCreate = api.Method{
		API:            Core_TaskScheduler,
		Version:        4,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	TaskDelete = api.Method{
		API:            Core_TaskScheduler,
		Version:        2,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	TaskUpdate = api.Method{
		API:            Core_TaskScheduler,
		Version:        4,
		Method:         api.MethodSet,
		ErrorSummaries: api.GlobalErrors,
	}
	TaskRun = api.Method{
		API:            Core_TaskScheduler,
		Version:        2,
		Method:         api.MethodRun,
		ErrorSummaries: api.GlobalErrors,
	}
	RootTaskCreate = api.Method{
		API:            Core_Task_Root,
		Version:        4,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	RootTaskUpdate = api.Method{
		API:            Core_Task_Root,
		Version:        4,
		Method:         api.MethodSet,
		ErrorSummaries: api.GlobalErrors,
	}
	RootTaskDelete = api.Method{
		API:            Core_Task_Root,
		Version:        4,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	PasswordConfirm = api.Method{
		API:            Core_Password_Confirm,
		Version:        2,
		Method:         api.MethodAuth,
		ErrorSummaries: api.GlobalErrors,
	}
)
