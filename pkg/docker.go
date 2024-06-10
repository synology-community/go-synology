package client

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api/docker"
	"github.com/synology-community/synology-api/pkg/api/docker/methods"
	"github.com/synology-community/synology-api/pkg/models"
)

type dockerClient struct {
	client *APIClient
}

// ProjectCleanStream implements docker.DockerApi.
func (d *dockerClient) ProjectCleanStream(ctx context.Context, req docker.ProjectCleanStreamRequest) (*docker.ProjectCleanStreamResponse, error) {
	return Get[docker.ProjectCleanStreamRequest, docker.ProjectCleanStreamResponse](d.client, ctx, &req, methods.ProjectCleanStream)
}

// ProjectCreate implements docker.DockerApi.
func (d *dockerClient) ProjectCreate(ctx context.Context, req docker.ProjectCreateRequest) (*docker.ProjectCreateResponse, error) {
	return Get[docker.ProjectCreateRequest, docker.ProjectCreateResponse](d.client, ctx, &req, methods.ProjectCreate)
}

// ProjectDelete implements docker.DockerApi.
func (d *dockerClient) ProjectDelete(ctx context.Context, req docker.ProjectDeleteRequest) (*docker.ProjectDeleteResponse, error) {
	return Get[docker.ProjectDeleteRequest, docker.ProjectDeleteResponse](d.client, ctx, &req, methods.ProjectDelete)
}

// ProjectGet implements docker.DockerApi.
func (d *dockerClient) ProjectGet(ctx context.Context, req docker.ProjectGetRequest) (*docker.ProjectGetResponse, error) {
	return Get[docker.ProjectGetRequest, docker.ProjectGetResponse](d.client, ctx, &req, methods.ProjectGet)
}

// ProjectList implements docker.DockerApi.
func (d *dockerClient) ProjectList(ctx context.Context, req docker.ProjectListRequest) (*docker.ProjectListResponse, error) {
	return Get[docker.ProjectListRequest, docker.ProjectListResponse](d.client, ctx, &req, methods.ProjectList)
}

// ProjectUpdate implements docker.DockerApi.
func (d *dockerClient) ProjectUpdate(ctx context.Context, req docker.ProjectUpdateRequest) (*docker.ProjectUpdateResponse, error) {
	return Get[docker.ProjectUpdateRequest, docker.ProjectUpdateResponse](d.client, ctx, &req, methods.ProjectUpdate)
}

// ImageDelete implements docker.DockerApi.
func (d *dockerClient) ImageDelete(ctx context.Context, req docker.ImageDeleteRequest) (*docker.ImageDeleteResponse, error) {
	return Get[docker.ImageDeleteRequest, docker.ImageDeleteResponse](d.client, ctx, &req, methods.ImageDelete)
}

// ImagePull implements docker.DockerApi.
func (d *dockerClient) ImagePull(ctx context.Context, repository string, tag string) (*docker.ImagePullStatusResponse, error) {
	res, err := d.ImagePullStart(ctx, docker.ImagePullStartRequest{
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
		task, err := d.ImagePullStatus(ctx, docker.ImagePullStatusRequest{
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

// ImagePullStart implements docker.DockerApi.
func (d *dockerClient) ImagePullStart(ctx context.Context, req docker.ImagePullStartRequest) (*docker.ImagePullStartResponse, error) {
	return Get[docker.ImagePullStartRequest, docker.ImagePullStartResponse](d.client, ctx, &req, methods.ImagePullStart)
}

// ImagePullStatus implements docker.DockerApi.
func (d *dockerClient) ImagePullStatus(ctx context.Context, req docker.ImagePullStatusRequest) (*docker.ImagePullStatusResponse, error) {
	return Get[docker.ImagePullStatusRequest, docker.ImagePullStatusResponse](d.client, ctx, &req, methods.ImagePullStatus)
}

func NewDockerClient(client *APIClient) docker.DockerApi {
	return &dockerClient{client: client}
}

// ContainerCreate implements docker.DockerApi.
func (d *dockerClient) ContainerCreate(ctx context.Context, req docker.CreateContainerRequest) (*docker.CreateContainerResponse, error) {
	return Get[docker.CreateContainerRequest, docker.CreateContainerResponse](d.client, ctx, &req, methods.Create)
}

// RegistryList implements docker.DockerApi.
func (d *dockerClient) RegistryList(ctx context.Context, req docker.ListRegistryRequest) (*docker.ListRegistryResponse, error) {
	return Get[docker.ListRegistryRequest, docker.ListRegistryResponse](d.client, ctx, &req, methods.RegistryList)
}
