package docker

import (
	"context"

	"github.com/synology-community/synology-api/pkg/models"
)

type DockerApi interface {
	ContainerCreate(ctx context.Context, container CreateContainerRequest) (*models.FolderList, error)
}
