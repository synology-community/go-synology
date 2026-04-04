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
	API_Network      = "SYNO.Virtualization.API.Network"
	SET_Guest        = "SYNO.Virtualization.Guest"
	SET_Guest_Image  = "SYNO.Virtualization.Guest.Image"
)

// VMMErrors contains error codes specific to the Synology Virtual Machine Manager API.
var VMMErrors = api.ErrorSummary{
	401:  "Bad parameter",
	402:  "Operation failed",
	403:  "Name conflict",
	404:  "The number of iSCSI LUNs has reached the system limit",
	500:  "The cluster is frozen. More than half of the hosts are offline",
	501:  "The cluster is in incompatible mode. Please upgrade to a compatible DSM version",
	600:  "The cluster is not ready",
	601:  "The host is offline",
	700:  "The storage is invalid",
	900:  "Failed to set a host to a virtual machine",
	901:  "The virtual machine does not have a host",
	902:  "Failed to power on: insufficient CPU threads",
	903:  "Failed to power on: insufficient memory",
	904:  "The virtual machine is already online",
	905:  "MAC conflict",
	906:  "Failed to create VM: the selected image is not found",
	907:  "The virtual machine is offline",
	908:  "Failed to power on: insufficient CPU threads for reservation on the host",
	909:  "Failed to power on: no corresponding networking on the host",
	910:  "Only the VirtIO hard disk controller can be used to boot the VM remotely",
	911:  "Virtual machines with UEFI enabled cannot be powered on remotely",
	1000: "Cannot find task_id",
	1001: "Need Virtual Machine Manager Pro",
	1400: "The result of image creating is partial success",
	1600: "VM edited successfully, but errors occurred while reserving memory or CPU on HA hosts",
}

var VMMErrorSummaries = api.GlobalErrors.Combine(VMMErrors)

var (
	GuestGet = api.Method{
		API:            API_Guest,
		Version:        1,
		Method:         "get",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestList = api.Method{
		API:            API_Guest,
		Version:        1,
		Method:         "list",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestCreate = api.Method{
		API:            API_Guest,
		Version:        1,
		Method:         "create",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestSet = api.Method{
		API:            API_Guest,
		Version:        1,
		Method:         "set",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestDelete = api.Method{
		API:            API_Guest,
		Version:        1,
		Method:         "delete",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestUpdate = api.Method{
		API:            SET_Guest,
		Version:        1,
		Method:         "set",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestPowerOn = api.Method{
		API:            API_Guest_Action,
		Version:        1,
		Method:         "poweron",
		ErrorSummaries: VMMErrorSummaries,
	}
	GuestPowerOff = api.Method{
		API:            API_Guest_Action,
		Version:        1,
		Method:         "poweroff",
		ErrorSummaries: VMMErrorSummaries,
	}
	ImageList = api.Method{
		API:            API_Guest_Image,
		Version:        1,
		Method:         "list",
		ErrorSummaries: VMMErrorSummaries,
	}
	ImageCreate = api.Method{
		API:            API_Guest_Image,
		Version:        1,
		Method:         "create",
		ErrorSummaries: VMMErrorSummaries,
	}
	ImageUploadAndCreate = api.Method{
		API:            SET_Guest_Image,
		Version:        1,
		Method:         "upload_and_create",
		ErrorSummaries: VMMErrorSummaries,
	}
	ImageDelete = api.Method{
		API:            API_Guest_Image,
		Version:        1,
		Method:         "delete",
		ErrorSummaries: VMMErrorSummaries,
	}
	TaskGet = api.Method{
		API:            API_Task_Info,
		Version:        1,
		Method:         "get",
		ErrorSummaries: VMMErrorSummaries,
	}
	StorageList = api.Method{
		API:            API_Storage,
		Version:        1,
		Method:         "list",
		ErrorSummaries: VMMErrorSummaries,
	}
	NetworkList = api.Method{
		API:            API_Network,
		Version:        1,
		Method:         "list",
		ErrorSummaries: VMMErrorSummaries,
	}
)
