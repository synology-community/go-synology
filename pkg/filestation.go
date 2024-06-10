package client

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/api/filestation/methods"
	"github.com/synology-community/synology-api/pkg/models"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type fileStationClient struct {
	client *APIClient
}

func NewFileStationClient(client *APIClient) filestation.FileStationApi {
	return &fileStationClient{client: client}
}

// List implements filestation.FileStationApi.
func (f *fileStationClient) List(ctx context.Context, folderPath string) (*models.FileList, error) {
	return Get[models.FileListRequest, models.FileList](f.client, ctx, &models.FileListRequest{
		FolderPath: folderPath,
		Additional: []string{"real_path", "size", "owner", "time", "perm", "mount_point_type", "type", "fileid"},
	}, methods.List)
}

func (f *fileStationClient) Delete(ctx context.Context, paths []string, accurateProgress bool) (*filestation.DeleteStatusResponse, error) {
	// Start Delete the file
	rdel, err := f.client.FileStationAPI().DeleteStart(ctx, paths, true)
	if err != nil {
		return nil, fmt.Errorf("Unable to delete file, got error: %s", err)
	}
	return f.client.FileStationAPI().DeleteStatus(ctx, rdel.TaskID)
}

func (f *fileStationClient) DeleteStart(ctx context.Context, paths []string, accurateProgress bool) (*filestation.DeleteStartResponse, error) {
	method := methods.DeleteStart
	return Get[filestation.DeleteStartRequest, filestation.DeleteStartResponse](f.client, ctx, &filestation.DeleteStartRequest{
		Paths:            paths,
		AccurateProgress: accurateProgress,
	}, method)
}

func (f *fileStationClient) DeleteStatus(ctx context.Context, taskID string) (*filestation.DeleteStatusResponse, error) {
	return Get[filestation.DeleteStatusRequest, filestation.DeleteStatusResponse](f.client, ctx, &filestation.DeleteStatusRequest{
		TaskID: taskID,
	}, methods.DeleteStatus)
}

func (f *fileStationClient) MD5(ctx context.Context, path string) (*filestation.MD5StatusResponse, error) {
	rmd5, err := f.client.FileStationAPI().MD5Start(ctx, path)

	time.Sleep(5000 * time.Millisecond)

	if err != nil {
		return nil, fmt.Errorf("unable to get file md5, got error: %s", err)
	}

	return f.client.FileStationAPI().MD5Status(ctx, rmd5.TaskID)
}

func (f *fileStationClient) MD5Start(ctx context.Context, path string) (*filestation.MD5StartResponse, error) {
	return Get[filestation.MD5StartRequest, filestation.MD5StartResponse](f.client, ctx, &filestation.MD5StartRequest{
		Path: path,
	}, methods.MD5Start)
}

func (f *fileStationClient) MD5Status(ctx context.Context, taskID string) (*filestation.MD5StatusResponse, error) {
	return Get[filestation.MD5StatusRequest, filestation.MD5StatusResponse](f.client, ctx, &filestation.MD5StatusRequest{
		TaskID: filestation.UrlWrapString(taskID),
	}, methods.MD5Status)
}

// Download implements filestation.FileStationApi.
func (f *fileStationClient) Download(ctx context.Context, path string, mode string) (*form.File, error) {
	return Get[filestation.DownloadRequest, form.File](f.client, ctx, &filestation.DownloadRequest{
		Path: path,
		Mode: mode,
	}, methods.Download)
}

// Rename implements filestation.FileStationApi.
func (f *fileStationClient) Rename(ctx context.Context, path string, name string, newName string) (*models.FileList, error) {
	return Get[filestation.RenameRequest, models.FileList](f.client, ctx, &filestation.RenameRequest{
		Path:    path,
		Name:    name,
		NewName: newName,
	}, methods.Rename)
}

// CreateFolder implements filestation.FileStationApi.
func (f *fileStationClient) CreateFolder(ctx context.Context, paths []string, names []string, forceParent bool) (*models.FolderList, error) {
	return Get[filestation.CreateFolderRequest, models.FolderList](f.client, ctx, &filestation.CreateFolderRequest{
		Paths:       paths,
		Names:       names,
		ForceParent: forceParent,
	}, methods.CreateFolder)
}

// ListShares implements filestation.FileStationApi.
func (f *fileStationClient) ListShares(ctx context.Context) (*models.ShareList, error) {
	return Get[filestation.ListShareRequest, models.ShareList](f.client, ctx, &filestation.ListShareRequest{}, methods.ListShares)
}

// Upload implements filestation.FileStationApi.
func (f *fileStationClient) Upload(ctx context.Context, path string, file form.File, createParents bool, overwrite bool) (*filestation.UploadResponse, error) {
	return PostFile[filestation.UploadRequest, filestation.UploadResponse](f.client, ctx, &filestation.UploadRequest{
		Path:          path,
		File:          file,
		CreateParents: createParents,
		Overwrite:     overwrite,
	}, methods.Upload)
}
