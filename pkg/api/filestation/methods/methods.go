package methods

import (
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/filestation"
)

var (
	Upload = api.Method{
		API:     "SYNO.FileStation.Upload",
		Version: 2,
		Method:  "upload",
		ErrorSummary: filestation.CommonErrors.Combine(api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		}),
	}
	List = api.Method{
		API:          "SYNO.FileStation.List",
		Version:      2,
		Method:       "list",
		ErrorSummary: filestation.CommonErrors,
	}
	ListShares = api.Method{
		API:          "SYNO.FileStation.List",
		Version:      2,
		Method:       api.MethodListShares,
		ErrorSummary: filestation.CommonErrors,
	}
	Rename = api.Method{
		API:     "SYNO.FileStation.Rename",
		Version: 2,
		Method:  api.MethodRename,
		ErrorSummary: filestation.CommonErrors.Combine(api.ErrorSummary{
			1200: "Failed to rename it.",
		}),
	}
	Info = api.Method{
		API:          "SYNO.FileStation.Info",
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: filestation.CommonErrors,
	}
	CreateFolder = api.Method{
		API:     "SYNO.FileStation.CreateFolder",
		Version: 2,
		Method:  api.MethodCreate,
		ErrorSummary: api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		},
	}
	DeleteStart = api.Method{
		API:          "SYNO.FileStation.Delete",
		Version:      2,
		Method:       api.MethodStart,
		ErrorSummary: filestation.CommonErrors,
	}
	DeleteStatus = api.Method{
		API:          "SYNO.FileStation.Delete",
		Version:      1,
		Method:       api.MethodStatus,
		ErrorSummary: filestation.CommonErrors,
	}
	Download = api.Method{
		API:          "SYNO.FileStation.Download",
		Version:      2,
		Method:       "download",
		ErrorSummary: filestation.CommonErrors,
	}
	MD5Start = api.Method{
		API:          "SYNO.FileStation.MD5",
		Version:      2,
		Method:       api.MethodStatus,
		ErrorSummary: filestation.CommonErrors,
	}
	MD5Status = api.Method{
		API:          "SYNO.FileStation.MD5",
		Version:      2,
		Method:       "status",
		ErrorSummary: filestation.CommonErrors,
	}
	BackgroundTask = api.Method{
		API:     "SYNO.FileStation.BackgroundTask",
		Version: 3,
	}
	CheckExist = api.Method{
		API:     "SYNO.FileStation.CheckExist",
		Version: 2,
	}
	CheckPermission = api.Method{
		API:     "SYNO.FileStation.CheckPermission",
		Version: 3,
	}
	Compress = api.Method{
		API:     "SYNO.FileStation.Compress",
		Version: 3,
	}
	CopyMove = api.Method{
		API:     "SYNO.FileStation.CopyMove",
		Version: 3,
	}
	Delete = api.Method{
		API:     "SYNO.FileStation.Delete",
		Version: 2,
	}
	DirSize = api.Method{
		API:     "SYNO.FileStation.DirSize",
		Version: 2,
	}
	ExternalGoogleDrive = api.Method{
		API:     "SYNO.FileStation.External.GoogleDrive",
		Version: 2,
	}
	Extract = api.Method{
		API:     "SYNO.FileStation.Extract",
		Version: 2,
	}
	Favorite = api.Method{
		API:     "SYNO.FileStation.Favorite",
		Version: 2,
	}
	FormUpload = api.Method{
		API:     "SYNO.FileStation.FormUpload",
		Version: 2,
	}
	MD5 = api.Method{
		API:     "SYNO.FileStation.MD5",
		Version: 2,
	}
	Mount = api.Method{
		API:     "SYNO.FileStation.Mount",
		Version: 1,
	}
	MountList = api.Method{
		API:     "SYNO.FileStation.Mount.List",
		Version: 1,
	}
	Notify = api.Method{
		API:     "SYNO.FileStation.Notify",
		Version: 1,
	}
	PhotoUpload = api.Method{
		API:     "SYNO.FileStation.PhotoUpload",
		Version: 3,
	}
	Property = api.Method{
		API:     "SYNO.FileStation.Property",
		Version: 1,
	}
	PropertyACLOwner = api.Method{
		API:     "SYNO.FileStation.Property.ACLOwner",
		Version: 1,
	}
	PropertyCompressSize = api.Method{
		API:     "SYNO.FileStation.Property.CompressSize",
		Version: 1,
	}
	PropertyMtime = api.Method{
		API:     "SYNO.FileStation.Property.Mtime",
		Version: 1,
	}
	Search = api.Method{
		API:     "SYNO.FileStation.Search",
		Version: 2,
	}
	SearchHistory = api.Method{
		API:     "SYNO.FileStation.Search.History",
		Version: 1,
	}
	Settings = api.Method{
		API:     "SYNO.FileStation.Settings",
		Version: 1,
	}
	Sharing = api.Method{
		API:     "SYNO.FileStation.Sharing",
		Version: 3,
	}
	SharingDownload = api.Method{
		API:     "SYNO.FileStation.Sharing.Download",
		Version: 1,
	}
	Snapshot = api.Method{
		API:     "SYNO.FileStation.Snapshot",
		Version: 2,
	}
	Thumb = api.Method{
		API:     "SYNO.FileStation.Thumb",
		Version: 3,
	}
	Timeout = api.Method{
		API:     "SYNO.FileStation.Timeout",
		Version: 1,
	}
	UIString = api.Method{
		API:     "SYNO.FileStation.UIString",
		Version: 1,
	}
	UserGrp = api.Method{
		API:     "SYNO.FileStation.UserGrp",
		Version: 1,
	}
	VFSConnection = api.Method{
		API:     "SYNO.FileStation.VFS.Connection",
		Version: 1,
	}
	VFSFile = api.Method{
		API:     "SYNO.FileStation.VFS.File",
		Version: 1,
	}
	VFSGDrive = api.Method{
		API:     "SYNO.FileStation.VFS.GDrive",
		Version: 1,
	}
	VFSProfile = api.Method{
		API:     "SYNO.FileStation.VFS.Profile",
		Version: 1,
	}
	VFSProtocol = api.Method{
		API:     "SYNO.FileStation.VFS.Protocol",
		Version: 1,
	}
	VFSUser = api.Method{
		API:     "SYNO.FileStation.VFS.User",
		Version: 1,
	}
	VirtualFolder = api.Method{
		API:     "SYNO.FileStation.VirtualFolder",
		Version: 2,
	}
	Worm = api.Method{
		API:     "SYNO.FileStation.Worm",
		Version: 2,
	}
	WormLock = api.Method{
		API:     "SYNO.FileStation.Worm.Lock",
		Version: 2,
	}
)
