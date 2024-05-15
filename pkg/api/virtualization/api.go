package virtualization

import (
	"context"
)

type VirtualizationAPI interface {
	GuestGet(ctx context.Context, guest Guest) (*Guest, error)
	GuestList(ctx context.Context) (*GuestList, error)
	GuestCreate(ctx context.Context, guest Guest) (*Guest, error)
	GuestDelete(ctx context.Context, guest Guest) error

	ImageList(ctx context.Context) (*ImageList, error)
	ImageCreate(ctx context.Context, image Image) (*Task, error)
	ImageDelete(ctx context.Context, imageID string) error

	TaskGet(ctx context.Context, taskID string) (*Task, error)

	StorageList(ctx context.Context) (*StorageList, error)
}
