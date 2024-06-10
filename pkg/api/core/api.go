package core

import (
	"context"
)

type CoreApi interface {
	PackageList(ctx context.Context) (*PackageListResponse, error)
}
