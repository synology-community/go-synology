package core

import (
	"context"
)

type Api interface {
	PackageList(ctx context.Context) (*PackageListResponse, error)
	PackageGet(ctx context.Context, id string) (*PackageGetResponse, error)
	SystemInfo(ctx context.Context) (*SystemInfoResponse, error)
	PackageServerList(ctx context.Context, req PackageServerListRequest) (*PackageServerListResponse, error)
	PackageInstall(ctx context.Context, req PackageInstallRequest) (*PackageInstallResponse, error)
	PackageInstallStatus(ctx context.Context, req PackageInstallStatusRequest) (*PackageInstallStatusResponse, error)
	ContentLength(ctx context.Context, url string) (int64, error)
}
