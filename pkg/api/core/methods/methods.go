package methods

import (
	"github.com/synology-community/go-synology/pkg/api"
)

const (
	Core_Event                  = "SYNO.Core.EventScheduler"
	Core_Event_Root             = "SYNO.Core.EventScheduler.Root"
	Core_Network                = "SYNO.Core.Network"
	Core_Package                = "SYNO.Core.Package"
	Core_Package_Feed           = "SYNO.Core.Package.Feed"
	Core_Package_Installation   = "SYNO.Core.Package.Installation"
	Core_Package_Server         = "SYNO.Core.Package.Server"
	Core_Package_Setting        = "SYNO.Core.Package.Setting"
	Core_Package_Uninstallation = "SYNO.Core.Package.Uninstallation"
	Core_Password_Confirm       = "SYNO.Core.User.PasswordConfirm"
	Core_Share                  = "SYNO.Core.Share"
	Core_Storage_Volume         = "SYNO.Core.Storage.Volume"
	Core_System                 = "SYNO.Core.System"
	Core_Task_Root              = "SYNO.Core.TaskScheduler.Root"
	Core_TaskScheduler          = "SYNO.Core.TaskScheduler"
	Core_User                   = "SYNO.Core.User"
	DSM_PortEnable              = "SYNO.DSM.PortEnable"
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
	PackageInstallationUpload = api.Method{
		API:            Core_Package_Installation,
		Version:        1,
		Method:         api.MethodUpload,
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
		Method:         api.MethodIsPkgEnable,
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
	EventCreate = api.Method{
		API:            Core_Event,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	EventUpdate = api.Method{
		API:            Core_Event,
		Version:        1,
		Method:         api.MethodSet,
		ErrorSummaries: api.GlobalErrors,
	}
	EventDelete = api.Method{
		API:            Core_Event,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	EventRun = api.Method{
		API:            Core_Event,
		Version:        1,
		Method:         api.MethodRun,
		ErrorSummaries: api.GlobalErrors,
	}
	EventGet = api.Method{
		API:            Core_Event,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
	RootEventCreate = api.Method{
		API:            Core_Event_Root,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	RootEventUpdate = api.Method{
		API:            Core_Event_Root,
		Version:        1,
		Method:         api.MethodSet,
		ErrorSummaries: api.GlobalErrors,
	}
	RootEventDelete = api.Method{
		API:            Core_Event,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	PasswordConfirm = api.Method{
		API:            Core_Password_Confirm,
		Version:        2,
		Method:         api.MethodAuth,
		ErrorSummaries: api.GlobalErrors,
	}
	UserCreate = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	UserModify = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodSet,
		ErrorSummaries: api.GlobalErrors,
	}
	UserDelete = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	UserList = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	GroupCreate = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: api.GlobalErrors,
	}
	GroupModify = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodSet,
		ErrorSummaries: api.GlobalErrors,
	}
	GroupDelete = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: api.GlobalErrors,
	}
	GroupList = api.Method{
		API:            Core_User,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: api.GlobalErrors,
	}
	NetworkGet = api.Method{
		API:            Core_Network,
		Version:        2,
		Method:         api.MethodGet,
		ErrorSummaries: api.GlobalErrors,
	}
)
