package client

import (
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/core"
	"github.com/synology-community/synology-api/pkg/api/docker"
	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/api/virtualization"
)

type AuthStorage struct {
	SessionID string `url:"_sid"`
	Token     string `url:"SynoToken"`
}

type SynologyClient interface {
	api.Api

	VirtualizationAPI() virtualization.Api

	FileStationAPI() filestation.Api

	DockerAPI() docker.Api

	CoreAPI() core.Api

	// get(request api.Request, response api.Response) error
}
type APIClient struct {
	api.Api

	FileStation    filestation.Api
	Virtualization virtualization.Api
	Docker         docker.Api
	Core           core.Api
}

// FileStationAPI implements SynologyClient.
func (c *APIClient) FileStationAPI() filestation.Api {
	return c.FileStation
}

func (c *APIClient) VirtualizationAPI() virtualization.Api {
	return c.Virtualization
}

func (c *APIClient) DockerAPI() docker.Api {
	return c.Docker
}

func (c *APIClient) CoreAPI() core.Api {
	return c.Core
}

// New initializes "client" instance with minimal input configuration.
func New(o api.Options) (SynologyClient, error) {
	client, err := api.New(o)
	if err != nil {
		return nil, err
	}

	synoClient := &APIClient{
		Api: client,
	}

	synoClient.Core = core.New(synoClient)
	synoClient.FileStation = filestation.New(synoClient)
	synoClient.Virtualization = virtualization.New(synoClient)
	synoClient.Docker = docker.New(synoClient)

	return synoClient, nil
}
