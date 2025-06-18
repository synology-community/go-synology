package filestation

import (
	"context"
	"fmt"
	"path/filepath"
	"slices"
	"time"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/filestation/methods"
	"github.com/synology-community/go-synology/pkg/models"
	"github.com/synology-community/go-synology/pkg/util/form"
)

type Client struct {
	client api.Api
}

func New(client api.Api) Api {
	return &Client{client: client}
}

type FileNotFoundError struct {
	Path string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found: %s", e.Path)
}

// List implements FileStationApi.
func (f *Client) List(ctx context.Context, folderPath string) (*models.FileList, error) {
	return api.Get[models.FileList](f.client, ctx, &models.FileListRequest{
		FolderPath: folderPath,
		Additional: []string{
			"real_path",
			"size",
			"owner",
			"time",
			"perm",
			"mount_point_type",
			"type",
			"fileid",
		},
		FileType: "all",
	}, methods.List)
}

func (f *Client) Get(ctx context.Context, path string) (*models.File, error) {
	folder := filepath.Dir(path)
	resp, err := f.List(ctx, folder)
	if err != nil {
		return nil, fmt.Errorf("unable to get file, got error: %s", err)
	}
	if resp.Files == nil {
		return nil, fmt.Errorf("files is nil")
	}
	if len(resp.Files) == 0 {
		return nil, fmt.Errorf("result is empty")
	}
	i := slices.IndexFunc(resp.Files, func(f models.File) bool {
		return f.Path == path
	})
	if i == -1 {
		return nil, FileNotFoundError{Path: path}
	}
	return &resp.Files[i], nil
}

func (f *Client) Delete(
	ctx context.Context,
	paths []string,
	accurateProgress bool,
) (*DeleteStatusResponse, error) {
	// Start Delete the file
	rdel, err := f.DeleteStart(ctx, paths, true)
	if err != nil {
		return nil, fmt.Errorf("unable to delete file, got error: %s", err)
	}
	return f.DeleteStatus(ctx, rdel.TaskID)
}

func (f *Client) DeleteStart(
	ctx context.Context,
	paths []string,
	accurateProgress bool,
) (*DeleteStartResponse, error) {
	method := methods.DeleteStart
	return api.Get[DeleteStartResponse](f.client, ctx, &DeleteStartRequest{
		Paths:            paths,
		AccurateProgress: accurateProgress,
	}, method)
}

func (f *Client) DeleteStatus(ctx context.Context, taskID string) (*DeleteStatusResponse, error) {
	return api.Get[DeleteStatusResponse](f.client, ctx, &DeleteStatusRequest{
		TaskID: taskID,
	}, methods.DeleteStatus)
}

func (f *Client) MD5(ctx context.Context, path string) (*MD5StatusResponse, error) {
	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(60 * time.Second)
	}

	rmd5, err := f.MD5Start(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("unable to get file md5, got error: %s", err)
	}

	delay := 1 * time.Second
	for {
		task, err := f.MD5Status(ctx, rmd5.TaskID)
		if err != nil && task == nil {
			return nil, err
		}
		if task.Finished {
			return task, nil
		}
		if time.Now().After(deadline.Add(delay)) {
			return nil, fmt.Errorf("timeout waiting for task to complete")
		}
		time.Sleep(delay)
	}
}

func (f *Client) MD5Start(ctx context.Context, path string) (*MD5StartResponse, error) {
	return api.Get[MD5StartResponse](f.client, ctx, &MD5StartRequest{
		Path: path,
	}, methods.MD5Start)
}

func (f *Client) MD5Status(ctx context.Context, taskID string) (*MD5StatusResponse, error) {
	return api.Get[MD5StatusResponse](f.client, ctx, &MD5StatusRequest{
		TaskID: taskID,
	}, methods.MD5Status)
}

// Download implements FileStationApi.
func (f *Client) Download(ctx context.Context, path string, mode string) (*form.File, error) {
	return api.Get[form.File](f.client, ctx, &DownloadRequest{
		Path: path,
		Mode: mode,
	}, methods.Download)
}

// Rename implements FileStationApi.
func (f *Client) Rename(
	ctx context.Context,
	path string,
	name string,
	newName string,
) (*models.FileList, error) {
	return api.Get[models.FileList](f.client, ctx, &RenameRequest{
		Path:    path,
		Name:    name,
		NewName: newName,
	}, methods.Rename)
}

// CreateFolder implements FileStationApi.
func (f *Client) CreateFolder(
	ctx context.Context,
	paths []string,
	names []string,
	forceParent bool,
) (*models.FolderList, error) {
	return api.Get[models.FolderList](f.client, ctx, &CreateFolderRequest{
		Paths:       paths,
		Names:       names,
		ForceParent: forceParent,
	}, methods.CreateFolder)
}

// ListShares implements FileStationApi.
func (f *Client) ListShares(ctx context.Context) (*models.ShareList, error) {
	return api.Get[models.ShareList](f.client, ctx, &ListShareRequest{}, methods.ListShares)
}

// Upload implements FileStationApi.
func (f *Client) Upload(
	ctx context.Context,
	path string,
	file form.File,
	createParents bool,
	overwrite bool,
) (*UploadResponse, error) {
	return api.PostFile[UploadResponse](f.client, ctx, &UploadRequest{
		Path:          path,
		File:          file,
		CreateParents: createParents,
		Overwrite:     overwrite,
	}, methods.Upload)
}
