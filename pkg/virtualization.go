package client

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api/virtualization"
	"github.com/synology-community/synology-api/pkg/api/virtualization/methods"
	"github.com/synology-community/synology-api/pkg/models"
)

type virtualizationClient struct {
	client *APIClient
}

func (v *virtualizationClient) GuestPowerOn(ctx context.Context, guest virtualization.Guest) error {
	_, err := Get[virtualization.Guest, models.NilResponse](v.client, ctx, &guest, methods.GuestPowerOn)
	return err
}

// GuestPowerOff implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GuestPowerOff(ctx context.Context, guest virtualization.Guest) error {
	_, err := Get[virtualization.Guest, models.NilResponse](v.client, ctx, &guest, methods.GuestPowerOff)
	return err
}

// GuestUpdate implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GuestUpdate(ctx context.Context, guest virtualization.GuestUpdate) error {
	_, err := Post[virtualization.GuestUpdate, virtualization.GuestUpdateResponse](v.client, ctx, &guest, methods.GuestUpdate)
	return err
}

// StorageList implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) StorageList(ctx context.Context) (*virtualization.StorageList, error) {
	return List[virtualization.StorageList](v.client, ctx, methods.StorageList)
}

// ImageCreate implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ImageCreate(ctx context.Context, image virtualization.Image) (*virtualization.Task, error) {
	resp, err := Get[virtualization.Image, virtualization.TaskRef](v.client, ctx, &image, methods.ImageCreate)

	if err != nil {
		return nil, err
	}

	return v.TaskGet(ctx, resp.TaskID)
}

// ImageDelete implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ImageDelete(ctx context.Context, imageName string) error {
	_, err := Get[virtualization.Image, virtualization.TaskRef](v.client, ctx, &virtualization.Image{
		Name: imageName,
	}, methods.ImageDelete)

	return err
}

// ImageList implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ImageList(ctx context.Context) (*virtualization.ImageList, error) {
	return List[virtualization.ImageList](v.client, ctx, methods.ImageList)
}

// TaskGet implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) TaskGet(ctx context.Context, taskID string) (*virtualization.Task, error) {

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(60 * time.Second)
	}

	delay := 5 * time.Second
	for {
		task, err := Get[virtualization.TaskRef, virtualization.Task](v.client, ctx, &virtualization.TaskRef{
			TaskID: taskID,
		}, methods.TaskGet)
		if err != nil && task == nil {
			return nil, err
		}
		if task.Finished {
			return task, nil
		}
		if time.Now().After(deadline.Add(delay)) {
			return nil, fmt.Errorf("Timeout waiting for task to complete")
		}
		time.Sleep(delay)
	}
}

// GetGuest implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GuestGet(ctx context.Context, guest virtualization.Guest) (*virtualization.Guest, error) {
	return Get[virtualization.GetGuest, virtualization.Guest](v.client, ctx, &virtualization.GetGuest{Name: guest.Name}, methods.GuestGet)
}

// ListGuests implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GuestList(ctx context.Context) (*virtualization.GuestList, error) {
	return List[virtualization.GuestList](v.client, ctx, methods.GuestList)
}

// GuestCreate implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GuestCreate(ctx context.Context, guest virtualization.Guest) (*virtualization.Guest, error) {
	resp, err := Get[virtualization.Guest, virtualization.TaskRef](v.client, ctx, &guest, methods.GuestCreate)
	if err != nil {
		return nil, err
	}

	task, err := v.TaskGet(ctx, resp.TaskID)
	if err != nil {
		return nil, err
	}

	guest.ID = task.TaskInfo.GuestID

	return &guest, nil
}

// GuestDelete implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) GuestDelete(ctx context.Context, guest virtualization.Guest) error {
	_, err := Get[virtualization.Guest, virtualization.TaskRef](v.client, ctx, &virtualization.Guest{
		Name: guest.Name,
	}, methods.GuestDelete)

	return err
}

func NewVirtualizationClient(client *APIClient) virtualization.VirtualizationAPI {
	return &virtualizationClient{client: client}
}
