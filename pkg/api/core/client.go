package core

import (
	"context"

	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/core/methods"
)

type Client struct {
	client api.Api
}

// SystemInfo implements CoreApi.
func (c Client) SystemInfo(ctx context.Context) (*SystemInfoResponse, error) {
	panic("unimplemented")
}

// PackageList implements CoreApi.
func (c Client) PackageList(ctx context.Context) (*PackageListResponse, error) {
	return api.List[PackageListResponse](c.client, ctx, methods.PackageList)
}

func New(client api.Api) Api {
	return &Client{client: client}
}
