package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/docker/methods"
	"github.com/synology-community/synology-api/pkg/models"
)

type Client struct {
	client api.Api
}

// ProjectCleanStream implements DockerApi.
func (d *Client) ProjectCleanStream(ctx context.Context, req ProjectCleanStreamRequest) (*ProjectCleanStreamResponse, error) {
	return api.Get[ProjectCleanStreamRequest, ProjectCleanStreamResponse](d.client, ctx, &req, methods.ProjectCleanStream)
}

// ProjectCreate implements DockerApi.
func (d *Client) ProjectCreate(ctx context.Context, req ProjectCreateRequest) (*ProjectCreateResponse, error) {
	return api.Get[ProjectCreateRequest, ProjectCreateResponse](d.client, ctx, &req, methods.ProjectCreate)
}

// ProjectDelete implements DockerApi.
func (d *Client) ProjectDelete(ctx context.Context, req ProjectDeleteRequest) (*ProjectDeleteResponse, error) {
	return api.Get[ProjectDeleteRequest, ProjectDeleteResponse](d.client, ctx, &req, methods.ProjectDelete)
}

// ProjectGet implements DockerApi.
func (d *Client) ProjectGet(ctx context.Context, req ProjectGetRequest) (*ProjectGetResponse, error) {
	return api.Get[ProjectGetRequest, ProjectGetResponse](d.client, ctx, &req, methods.ProjectGet)
}

// ProjectList implements DockerApi.
func (d *Client) ProjectList(ctx context.Context, req ProjectListRequest) (*ProjectListResponse, error) {
	return api.Get[ProjectListRequest, ProjectListResponse](d.client, ctx, &req, methods.ProjectList)
}

// ProjectUpdate implements DockerApi.
func (d *Client) ProjectUpdate(ctx context.Context, req ProjectUpdateRequest) (*ProjectUpdateResponse, error) {
	return api.Get[ProjectUpdateRequest, ProjectUpdateResponse](d.client, ctx, &req, methods.ProjectUpdate)
}

// ImageDelete implements DockerApi.
func (d *Client) ImageDelete(ctx context.Context, req ImageDeleteRequest) (*ImageDeleteResponse, error) {
	return api.Get[ImageDeleteRequest, ImageDeleteResponse](d.client, ctx, &req, methods.ImageDelete)
}

// ImagePull implements DockerApi.
func (d *Client) ImagePull(ctx context.Context, repository string, tag string) (*ImagePullStatusResponse, error) {
	res, err := d.ImagePullStart(ctx, ImagePullStartRequest{
		Repository: models.JsonString(repository),
		Tag:        models.JsonString(tag),
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to delete file, got error: %s", err)
	}

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(60 * time.Second)
	}

	delay := 5 * time.Second
	for {
		task, err := d.ImagePullStatus(ctx, ImagePullStatusRequest{
			TaskID: models.JsonString(res.TaskID),
		})
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

// ImagePullStart implements DockerApi.
func (d *Client) ImagePullStart(ctx context.Context, req ImagePullStartRequest) (*ImagePullStartResponse, error) {
	return api.Get[ImagePullStartRequest, ImagePullStartResponse](d.client, ctx, &req, methods.ImagePullStart)
}

// ImagePullStatus implements DockerApi.
func (d *Client) ImagePullStatus(ctx context.Context, req ImagePullStatusRequest) (*ImagePullStatusResponse, error) {
	return api.Get[ImagePullStatusRequest, ImagePullStatusResponse](d.client, ctx, &req, methods.ImagePullStatus)
}

// ContainerCreate implements DockerApi.
func (d *Client) ContainerCreate(ctx context.Context, req CreateContainerRequest) (*CreateContainerResponse, error) {
	return api.Get[CreateContainerRequest, CreateContainerResponse](d.client, ctx, &req, methods.Create)
}

// RegistryList implements DockerApi.
func (d *Client) RegistryList(ctx context.Context, req ListRegistryRequest) (*ListRegistryResponse, error) {
	return api.Get[ListRegistryRequest, ListRegistryResponse](d.client, ctx, &req, methods.RegistryList)
}

func New(client api.Api) Api {
	return &Client{client: client}
}
