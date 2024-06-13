package core

import (
	"context"
)

type Api interface {
	PackageList(ctx context.Context) (*PackageListResponse, error)
	PackageGet(ctx context.Context, id string) (*PackageGetResponse, error)
	PackageFind(ctx context.Context, name string) (*Package, error)
	SystemInfo(ctx context.Context) (*SystemInfoResponse, error)
	PackageServerList(ctx context.Context, req PackageServerListRequest) (*PackageServerListResponse, error)
	PackageInstall(ctx context.Context, req PackageInstallRequest) (*PackageInstallResponse, error)
	PackageInstallCompound(ctx context.Context, name string, url string, size int64) error
	PackageInstallStatus(ctx context.Context, req PackageInstallStatusRequest) (*PackageInstallStatusResponse, error)
	PackageInstallDelete(ctx context.Context, req PackageInstallDeleteRequest) error
	PackageUninstall(ctx context.Context, req PackageUninstallRequest) (*PackageUninstallResponse, error)
	PackageUninstallCompound(ctx context.Context, name string) error
	ContentLength(ctx context.Context, url string) (int64, error)
	PackageFeedList(ctx context.Context) (*PackageFeedListResponse, error)
	PackageFeedAdd(ctx context.Context, req PackageFeedAddRequest) error
	PackageFeedDelete(ctx context.Context, req PackageFeedDeleteRequest) error
	PackageSettingGet(ctx context.Context, req PackageSettingGetRequest) (*PackageSettingGetResponse, error)
}
