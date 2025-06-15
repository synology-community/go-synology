package virtualization

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/virtualization/methods"
	"github.com/synology-community/go-synology/pkg/util/form"
)

type Client struct {
	client api.Api
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ImageUploadAndCreate implements Api.
func (v *Client) ImageUploadAndCreate(
	ctx context.Context,
	file form.File,
	imageRepos []string,
	imageType string,
) (*Task, error) {
	name := strings.TrimSuffix(path.Base(file.Name), path.Ext(file.Name))

	resp, err := api.PostFileWithQuery[TaskRef](v.client, ctx, &UploadAndCreateRequest{
		Name:       name,
		DsFilePath: path.Base(file.Name),
		ImageRepos: imageRepos,
		Type:       imageType,
		GetPatchBy: "upload",
		File:       file,
	}, methods.ImageUploadAndCreate)
	if err != nil {
		return nil, err
	}

	return v.TaskGet(ctx, resp.TaskID)
}

func (v *Client) GuestPowerOn(ctx context.Context, guest Guest) error {
	return api.Void(v.client, ctx, &guest, methods.GuestPowerOn)
}

// GuestPowerOff implements VirtualizationAPI.
func (v *Client) GuestPowerOff(ctx context.Context, guest Guest) error {
	return api.Void(v.client, ctx, &guest, methods.GuestPowerOff)
}

// GuestUpdate implements VirtualizationAPI.
func (v *Client) GuestUpdate(ctx context.Context, guest GuestUpdate) error {
	_, err := api.Post[api.Response](v.client, ctx, &guest, methods.GuestUpdate)
	return err
}

// StorageList implements VirtualizationAPI.
func (v *Client) StorageList(ctx context.Context) (*StorageList, error) {
	return api.List[StorageList](v.client, ctx, methods.StorageList)
}

// ImageCreate implements VirtualizationAPI.
func (v *Client) ImageCreate(ctx context.Context, image Image) (*Task, error) {
	resp, err := api.Post[TaskRef](v.client, ctx, &image, methods.ImageCreate)
	if err != nil {
		return nil, err
	}

	return v.TaskGet(ctx, resp.TaskID)
}

// ImageDelete implements VirtualizationAPI.
func (v *Client) ImageDelete(ctx context.Context, imageName string) error {
	_, err := api.Post[TaskRef](v.client, ctx, &Image{
		Name: imageName,
	}, methods.ImageDelete)

	return err
}

// ImageList implements VirtualizationAPI.
func (v *Client) ImageList(ctx context.Context) (*ImageList, error) {
	return api.Post[ImageList, api.Request](v.client, ctx, nil, methods.ImageList)
}

// TaskGet implements VirtualizationAPI.
func (v *Client) TaskGet(ctx context.Context, taskID string) (*Task, error) {
	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(60 * time.Second)
	}

	delay := 5 * time.Second
	for {
		task, err := api.Post[Task](v.client, ctx, &TaskRef{
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

// GetGuest implements VirtualizationAPI.
func (v *Client) GuestGet(ctx context.Context, guest Guest) (*Guest, error) {
	return api.Post[Guest](v.client, ctx, &GetGuest{Name: guest.Name}, methods.GuestGet)
}

// ListGuests implements VirtualizationAPI.
func (v *Client) GuestList(ctx context.Context) (*GuestList, error) {
	return api.Post[GuestList, api.Request](v.client, ctx, nil, methods.GuestList)
}

// GuestCreate implements VirtualizationAPI.
func (v *Client) GuestCreate(ctx context.Context, guest Guest) (*Guest, error) {
	resp, err := api.Post[TaskRef](v.client, ctx, &guest, methods.GuestCreate)
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

// GuestDelete implements VirtualizationAPI.
func (v *Client) GuestDelete(ctx context.Context, guest Guest) error {
	_, err := api.Post[TaskRef](v.client, ctx, &Guest{
		Name: guest.Name,
	}, methods.GuestDelete)

	return err
}

func New(client api.Api) Api {
	return &Client{client: client}
}
