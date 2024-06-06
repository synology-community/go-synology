package client

import (
	"context"

	"github.com/synology-community/synology-api/pkg/api/docker"
	"github.com/synology-community/synology-api/pkg/api/docker/methods"
	"github.com/synology-community/synology-api/pkg/models"
)

type dockerClient struct {
	client *APIClient
}

func NewDockerClient(client *APIClient) docker.DockerApi {
	return &dockerClient{client: client}
}

// ContainerCreate implements docker.DockerApi.
func (d *dockerClient) ContainerCreate(ctx context.Context, req docker.CreateContainerRequest) (*models.FolderList, error) {
	return Get[docker.CreateContainerRequest, models.FolderList](d.client, ctx, &req, methods.Create)
}
