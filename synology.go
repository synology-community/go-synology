package synology

import (
	"context"
	"errors"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/core"
	"github.com/synology-community/go-synology/pkg/api/docker"
	"github.com/synology-community/go-synology/pkg/api/filestation"
	"github.com/synology-community/go-synology/pkg/api/virtualization"
)

type AuthStorage struct {
	SessionID string `url:"_sid"`
	Token     string `url:"SynoToken"`
}

type Api interface {
	api.Api

	VirtualizationAPI() virtualization.Api

	FileStationAPI() filestation.Api

	DockerAPI() docker.Api

	CoreAPI() core.Api

	// get(request api.Request, response api.Response) error
}
type Client struct {
	api.Api

	FileStation    filestation.Api
	Virtualization virtualization.Api
	Docker         docker.Api
	Core           core.Api
}

// FileStationAPI implements SynologyClient.
func (c *Client) FileStationAPI() filestation.Api {
	return c.FileStation
}

func (c *Client) VirtualizationAPI() virtualization.Api {
	return c.Virtualization
}

func (c *Client) DockerAPI() docker.Api {
	return c.Docker
}

func (c *Client) CoreAPI() core.Api {
	return c.Core
}

// New initializes "client" instance with minimal input configuration.
func New(o api.Options) (Api, error) {
	client, err := api.New(o)
	if err != nil {
		return nil, err
	}

	synoClient := &Client{
		Api: client,
	}

	synoClient.Core = core.New(synoClient)
	synoClient.FileStation = filestation.New(synoClient)
	synoClient.Virtualization = virtualization.New(synoClient)
	synoClient.Docker = docker.New(synoClient)

	return synoClient, nil
}

// ExportSession exposes the underlying api.Client.ExportSession.
func (c *Client) ExportSession() api.Session {
	if ac, ok := c.Api.(*api.Client); ok {
		return ac.ExportSession()
	}
	return api.Session{}
}

// ImportSession exposes the underlying api.Client.ImportSession.
func (c *Client) ImportSession(s api.Session) {
	if ac, ok := c.Api.(*api.Client); ok {
		ac.ImportSession(s)
	}
}

// IsSessionAlive exposes the underlying api.Client.IsSessionAlive.
func (c *Client) IsSessionAlive(ctx context.Context) (bool, error) {
	if ac, ok := c.Api.(*api.Client); ok {
		return ac.IsSessionAlive(ctx)
	}
	return false, errors.New("underlying api client does not support IsSessionAlive")
}
