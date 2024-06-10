package core

import (
	"context"
)

type Api interface {
	PackageList(ctx context.Context) (*PackageListResponse, error)
	SystemInfo(ctx context.Context) (*SystemInfoResponse, error)
}
