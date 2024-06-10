package core

import (
	"context"
)

type CoreApi interface {
	PackageList(ctx context.Context) (*PackageListResponse, error)
	SystemInfo(ctx context.Context) (*SystemInfoResponse, error)
}
