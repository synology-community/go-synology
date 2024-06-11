package core

import (
	"context"

	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/core/methods"
	"github.com/synology-community/synology-api/pkg/models"
)

type Client struct {
	client api.Api
}

// ContentLength implements Api.
func (c *Client) ContentLength(ctx context.Context, url string) (int64, error) {
	resp, err := c.client.Client().Head(url)
	if err != nil {
		return int64(0), err
	}
	return resp.ContentLength, nil
}

// SystemInfo implements CoreApi.
func (c Client) SystemInfo(ctx context.Context) (*SystemInfoResponse, error) {
	panic("unimplemented")
}

// PackageList implements CoreApi.
func (c Client) PackageList(ctx context.Context) (*PackageListResponse, error) {
	return api.List[PackageListResponse](c.client, ctx, methods.PackageList)
}

func (c Client) PackageServerList(ctx context.Context, req PackageServerListRequest) (*PackageServerListResponse, error) {
	return api.Get[PackageServerListRequest, PackageServerListResponse](c.client, ctx, &req, methods.PackageServerList)
}

func (c Client) PackageGet(ctx context.Context, id string) (*PackageGetResponse, error) {
	return api.Get[PackageGetRequest, PackageGetResponse](c.client, ctx, &PackageGetRequest{ID: models.JsonString(id)}, methods.PackageGet)
}

func (c Client) PackageInstallStatus(ctx context.Context, req PackageInstallStatusRequest) (*PackageInstallStatusResponse, error) {
	return api.Get[PackageInstallStatusRequest, PackageInstallStatusResponse](c.client, ctx, &req, methods.PackageInstallationStatus)
}

func (c Client) PackageInstall(ctx context.Context, req PackageInstallRequest) (*PackageInstallResponse, error) {
	if req.URL != "" && req.FileSize < 1 {
		size, err := c.ContentLength(ctx, req.URL)
		if err != nil {
			return nil, err
		}
		req.FileSize = size
	}

	return api.Get[PackageInstallRequest, PackageInstallResponse](c.client, ctx, &req, methods.PackageInstallationInstall)
}

func New(client api.Api) Api {
	return &Client{client: client}
}
