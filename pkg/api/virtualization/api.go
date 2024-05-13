package virtualization

import (
	"context"

	"github.com/synology-community/synology-api/pkg/api"
)

type VirtualizationAPI interface {
	GetGuest(ctx context.Context, name string) (*Guest, error)
	ListGuests(ctx context.Context) (*GuestList, error)

	ImageList(ctx context.Context) (*ImageList, error)
	ImageCreate(ctx context.Context, image Image) (*Task, error)
	ImageDelete(ctx context.Context, imageID string) (*TaskRef, error)

	TaskGet(ctx context.Context, taskID string) (*Task, error)

	StorageList(ctx context.Context) (*StorageList, error)
}

var API_METHODS = api.APIMethodLookup{
	"GetGuest": {
		API:          "SYNO.Virtualization.API.Guest",
		Version:      1,
		Method:       "get",
		ErrorSummary: api.GlobalErrors,
	},
	"ListGuests": {
		API:          "SYNO.Virtualization.API.Guest",
		Version:      1,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	},
	"ImageList": {
		API:          "SYNO.Virtualization.API.Guest.Image",
		Version:      2,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	},
	"ImageCreate": {
		API:          "SYNO.Virtualization.API.Guest.Image",
		Version:      1,
		Method:       "create",
		ErrorSummary: api.GlobalErrors,
	},
	"ImageDelete": {
		API:          "SYNO.Virtualization.API.Guest.Image",
		Version:      1,
		Method:       "delete",
		ErrorSummary: api.GlobalErrors,
	},
	"TaskGet": {
		API:          "SYNO.Virtualization.API.Task.Info",
		Version:      1,
		Method:       "get",
		ErrorSummary: api.GlobalErrors,
	},
	"StorageList": {
		API:          "SYNO.Virtualization.API.Storage",
		Version:      1,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	},
}
