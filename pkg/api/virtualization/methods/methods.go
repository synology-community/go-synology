package methods

import (
	"github.com/synology-community/go-synology/pkg/api"
)

const (
	API_Guest        = "SYNO.Virtualization.API.Guest"
	API_Guest_Image  = "SYNO.Virtualization.API.Guest.Image"
	API_Guest_Action = "SYNO.Virtualization.API.Guest.Action"
	API_Task_Info    = "SYNO.Virtualization.API.Task.Info"
	API_Storage      = "SYNO.Virtualization.API.Storage"
	SET_Guest        = "SYNO.Virtualization.Guest"
)

var (
	GuestGet = api.Method{
		API:          API_Guest,
		Version:      1,
		Method:       "get",
		ErrorSummary: api.GlobalErrors,
	}
	GuestList = api.Method{
		API:          API_Guest,
		Version:      1,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	}
	GuestCreate = api.Method{
		API:          API_Guest,
		Version:      1,
		Method:       "create",
		ErrorSummary: api.GlobalErrors,
	}
	GuestDelete = api.Method{
		API:          API_Guest,
		Version:      1,
		Method:       "delete",
		ErrorSummary: api.GlobalErrors,
	}
	GuestUpdate = api.Method{
		API:          SET_Guest,
		Version:      1,
		Method:       "set",
		ErrorSummary: api.GlobalErrors,
	}
	GuestPowerOn = api.Method{
		API:          API_Guest_Action,
		Version:      1,
		Method:       "poweron",
		ErrorSummary: api.GlobalErrors,
	}
	GuestPowerOff = api.Method{
		API:          API_Guest_Action,
		Version:      1,
		Method:       "poweroff",
		ErrorSummary: api.GlobalErrors,
	}
	ImageList = api.Method{
		API:          API_Guest_Image,
		Version:      1,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	}
	ImageCreate = api.Method{
		API:          API_Guest_Image,
		Version:      1,
		Method:       "create",
		ErrorSummary: api.GlobalErrors,
	}
	ImageDelete = api.Method{
		API:          API_Guest_Image,
		Version:      1,
		Method:       "delete",
		ErrorSummary: api.GlobalErrors,
	}
	TaskGet = api.Method{
		API:          API_Task_Info,
		Version:      1,
		Method:       "get",
		ErrorSummary: api.GlobalErrors,
	}
	StorageList = api.Method{
		API:          API_Storage,
		Version:      1,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	}
)
