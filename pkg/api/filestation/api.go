package filestation

import (
	"context"

	"github.com/synology-community/synology-api/pkg/models"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type Api interface {
	CreateFolder(ctx context.Context, paths []string, names []string, forceParent bool) (*models.FolderList, error)
	ListShares(ctx context.Context) (*models.ShareList, error)
	List(ctx context.Context, folderPath string) (*models.FileList, error)
	Upload(ctx context.Context, path string, file form.File, createParents bool, overwrite bool) (*UploadResponse, error)
	Rename(ctx context.Context, path string, name string, newName string) (*models.FileList, error)
	Download(ctx context.Context, path string, mode string) (*form.File, error)
	Delete(ctx context.Context, paths []string, accurateProgress bool) (*DeleteStatusResponse, error)
	DeleteStart(ctx context.Context, paths []string, accurateProgress bool) (*DeleteStartResponse, error)
	DeleteStatus(ctx context.Context, taskID string) (*DeleteStatusResponse, error)
	MD5(ctx context.Context, path string) (*MD5StatusResponse, error)
	MD5Start(ctx context.Context, path string) (*MD5StartResponse, error)
	MD5Status(ctx context.Context, taskID string) (*MD5StatusResponse, error)
}
