package core

import (
	"context"
	"fmt"
	"slices"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/core/methods"
	"github.com/synology-community/go-synology/pkg/models"
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

func (c Client) PackageCheck(ctx context.Context) (*Package, error) {
	panic("unimplemented")
}

// PackageList implements CoreApi.
func (c Client) PackageList(ctx context.Context) (*PackageListResponse, error) {
	return api.Get[PackageListRequest, PackageListResponse](c.client, ctx, &PackageListRequest{
		Additional: []string{"description", "description_enu", "dependent_packages", "beta", "distributor", "distributor_url", "maintainer", "maintainer_url", "dsm_apps", "dsm_app_page", "dsm_app_launch_name", "report_beta_url", "support_center", "startable", "installed_info", "support_url", "is_uninstall_pages", "install_type", "autoupdate", "silent_upgrade", "installing_progress", "ctl_uninstall", "updated_at", "status", "url", "available_operation", "install_type"},
	}, methods.PackageList)
}

func (c Client) PackageFind(ctx context.Context, name string) (*Package, error) {
	for i := 0; i < 2; i++ {
		var req PackageServerListRequest
		if i == 0 {
			req.LoadOthers = false
		} else {
			req.LoadOthers = true
		}
		resp, err := c.PackageServerList(ctx, req)
		if err != nil {
			return nil, err
		}

		ii := slices.IndexFunc(resp.Packages, func(p Package) bool {
			return p.Package == name
		})
		if ii != -1 {
			return &resp.Packages[ii], nil
		}
	}
	return nil, fmt.Errorf("package %s not found", name)
}

func (c Client) PackageServerList(ctx context.Context, req PackageServerListRequest) (*PackageServerListResponse, error) {
	return api.Get[PackageServerListRequest, PackageServerListResponse](c.client, ctx, &req, methods.PackageServerList)
}

func (c Client) PackageGet(ctx context.Context, id string) (*PackageGetResponse, error) {
	return api.Get[PackageGetRequest, PackageGetResponse](c.client, ctx, &PackageGetRequest{
		ID:         models.JsonString(id),
		Additional: []string{"description", "description_enu", "dependent_packages", "beta", "distributor", "distributor_url", "maintainer", "maintainer_url", "dsm_apps", "dsm_app_page", "dsm_app_launch_name", "report_beta_url", "support_center", "startable", "installed_info", "support_url", "is_uninstall_pages", "install_type", "autoupdate", "silent_upgrade", "installing_progress", "ctl_uninstall", "updated_at", "status", "url", "available_operation", "install_type"},
	}, methods.PackageGet)
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

	req.Operation = "install"

	return api.Get[PackageInstallRequest, PackageInstallResponse](c.client, ctx, &req, methods.PackageInstallationInstall)
}

func (c Client) PackageUninstall(ctx context.Context, req PackageUninstallRequest) (*PackageUninstallResponse, error) {
	return api.Get[PackageUninstallRequest, PackageUninstallResponse](c.client, ctx, &req, methods.PackageUnistallationUninstall)
}

func New(client api.Api) Api {
	return &Client{client: client}
}
