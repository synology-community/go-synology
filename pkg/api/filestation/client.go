package filestation

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/filestation/methods"
	"github.com/synology-community/synology-api/pkg/models"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type Client struct {
	client api.Api
}

func New(client api.Api) Api {
	return &Client{client: client}
}

// List implements FileStationApi.
func (f *Client) List(ctx context.Context, folderPath string) (*models.FileList, error) {
	return api.Get[models.FileListRequest, models.FileList](f.client, ctx, &models.FileListRequest{
		FolderPath: folderPath,
		Additional: []string{"real_path", "size", "owner", "time", "perm", "mount_point_type", "type", "fileid"},
	}, methods.List)
}

func (f *Client) Delete(ctx context.Context, paths []string, accurateProgress bool) (*DeleteStatusResponse, error) {
	// Start Delete the file
	rdel, err := f.DeleteStart(ctx, paths, true)
	if err != nil {
		return nil, fmt.Errorf("Unable to delete file, got error: %s", err)
	}
	return f.DeleteStatus(ctx, rdel.TaskID)
}

func (f *Client) DeleteStart(ctx context.Context, paths []string, accurateProgress bool) (*DeleteStartResponse, error) {
	method := methods.DeleteStart
	return api.Get[DeleteStartRequest, DeleteStartResponse](f.client, ctx, &DeleteStartRequest{
		Paths:            paths,
		AccurateProgress: accurateProgress,
	}, method)
}

func (f *Client) DeleteStatus(ctx context.Context, taskID string) (*DeleteStatusResponse, error) {
	return api.Get[DeleteStatusRequest, DeleteStatusResponse](f.client, ctx, &DeleteStatusRequest{
		TaskID: taskID,
	}, methods.DeleteStatus)
}

func (f *Client) MD5(ctx context.Context, path string) (*MD5StatusResponse, error) {
	rmd5, err := f.MD5Start(ctx, path)

	time.Sleep(5000 * time.Millisecond)

	if err != nil {
		return nil, fmt.Errorf("unable to get file md5, got error: %s", err)
	}

	return f.MD5Status(ctx, rmd5.TaskID)
}

func (f *Client) MD5Start(ctx context.Context, path string) (*MD5StartResponse, error) {
	return api.Get[MD5StartRequest, MD5StartResponse](f.client, ctx, &MD5StartRequest{
		Path: path,
	}, methods.MD5Start)
}

func (f *Client) MD5Status(ctx context.Context, taskID string) (*MD5StatusResponse, error) {
	return api.Get[MD5StatusRequest, MD5StatusResponse](f.client, ctx, &MD5StatusRequest{
		TaskID: UrlWrapString(taskID),
	}, methods.MD5Status)
}

// Download implements FileStationApi.
func (f *Client) Download(ctx context.Context, path string, mode string) (*form.File, error) {
	return api.Get[DownloadRequest, form.File](f.client, ctx, &DownloadRequest{
		Path: path,
		Mode: mode,
	}, methods.Download)
}

// Rename implements FileStationApi.
func (f *Client) Rename(ctx context.Context, path string, name string, newName string) (*models.FileList, error) {
	return api.Get[RenameRequest, models.FileList](f.client, ctx, &RenameRequest{
		Path:    path,
		Name:    name,
		NewName: newName,
	}, methods.Rename)
}

// CreateFolder implements FileStationApi.
func (f *Client) CreateFolder(ctx context.Context, paths []string, names []string, forceParent bool) (*models.FolderList, error) {
	return api.Get[CreateFolderRequest, models.FolderList](f.client, ctx, &CreateFolderRequest{
		Paths:       paths,
		Names:       names,
		ForceParent: forceParent,
	}, methods.CreateFolder)
}

// ListShares implements FileStationApi.
func (f *Client) ListShares(ctx context.Context) (*models.ShareList, error) {
	return api.Get[ListShareRequest, models.ShareList](f.client, ctx, &ListShareRequest{}, methods.ListShares)
}

// Upload implements FileStationApi.
func (f *Client) Upload(ctx context.Context, path string, file form.File, createParents bool, overwrite bool) (*UploadResponse, error) {
	return api.PostFile[UploadRequest, UploadResponse](f.client, ctx, &UploadRequest{
		Path:          path,
		File:          file,
		CreateParents: createParents,
		Overwrite:     overwrite,
	}, methods.Upload)
}
