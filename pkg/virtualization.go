package client

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api/virtualization"
)

type virtualizationClient struct {
	client *APIClient
}

// StorageList implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) StorageList(ctx context.Context) (*virtualization.StorageList, error) {
	return List[virtualization.StorageList](v.client, ctx, virtualization.API_METHODS["StorageList"])
}

// ImageCreate implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ImageCreate(ctx context.Context, image virtualization.Image) (*virtualization.Task, error) {
	resp, err := Get[virtualization.Image, virtualization.TaskRef](v.client, ctx, &image, virtualization.API_METHODS["ImageCreate"])

	if err != nil {
		return nil, err
	}

	return v.TaskGet(ctx, resp.TaskID)
}

// ImageDelete implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ImageDelete(ctx context.Context, imageID string) (*virtualization.TaskRef, error) {
	return Get[virtualization.Image, virtualization.TaskRef](v.client, ctx, &virtualization.Image{
		ID: imageID,
	}, virtualization.API_METHODS["ImageDelete"])
}

// ImageList implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ImageList(ctx context.Context) (*virtualization.ImageList, error) {
	return List[virtualization.ImageList](v.client, ctx, virtualization.API_METHODS["ImageList"])
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
		}, virtualization.API_METHODS["TaskGet"])
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
func (v *virtualizationClient) GetGuest(ctx context.Context, name string) (*virtualization.Guest, error) {
	return Get[virtualization.GetGuest, virtualization.Guest](v.client, ctx, &virtualization.GetGuest{Name: name}, virtualization.API_METHODS["GetGuest"])
}

// ListGuests implements virtualization.VirtualizationAPI.
func (v *virtualizationClient) ListGuests(ctx context.Context) (*virtualization.GuestList, error) {
	return List[virtualization.GuestList](v.client, ctx, virtualization.API_METHODS["ListGuests"])
}

func NewVirtualizationClient(client *APIClient) virtualization.VirtualizationAPI {
	return &virtualizationClient{client: client}
}
