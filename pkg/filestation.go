package client

import (
	"github.com/synology-community/synology-api/package/api/filestation"
	"github.com/synology-community/synology-api/package/models"
	"github.com/synology-community/synology-api/package/util/form"
)

type fileStationClient struct {
	client *APIClient
}

func (f *fileStationClient) DeleteStart(paths []string, accurateProgress bool) (*filestation.DeleteStartResponse, error) {
	return Get[filestation.DeleteStartRequest, filestation.DeleteStartResponse](f.client, &filestation.DeleteStartRequest{
		Paths:            paths,
		AccurateProgress: accurateProgress,
	})
}

func (f *fileStationClient) DeleteStatus(taskID string) (*filestation.DeleteStatusResponse, error) {
	return Get[filestation.DeleteStatusRequest, filestation.DeleteStatusResponse](f.client, &filestation.DeleteStatusRequest{
		TaskID: taskID,
	})
}

func (f *fileStationClient) MD5Start(path string) (*filestation.MD5StartResponse, error) {
	return Get[filestation.MD5StartRequest, filestation.MD5StartResponse](f.client, &filestation.MD5StartRequest{
		Path: path,
	})
}

func (f *fileStationClient) MD5Status(taskID string) (*filestation.MD5StatusResponse, error) {
	return Get[filestation.MD5StatusRequest, filestation.MD5StatusResponse](f.client, &filestation.MD5StatusRequest{
		TaskID: taskID,
	})
}

// Download implements filestation.FileStationApi.
func (f *fileStationClient) Download(path string, mode string) (*filestation.DownloadResponse, error) {
	return Get[filestation.DownloadRequest, filestation.DownloadResponse](f.client, &filestation.DownloadRequest{
		Path: path,
		Mode: mode,
	})
}

// Rename implements filestation.FileStationApi.
func (f *fileStationClient) Rename(path string, name string, newName string) (*models.FileList, error) {
	return Get[filestation.RenameRequest, models.FileList](f.client, &filestation.RenameRequest{
		Path:    path,
		Name:    name,
		NewName: newName,
	})
}

// CreateFolder implements filestation.FileStationApi.
func (f *fileStationClient) CreateFolder(paths []string, names []string, forceParent bool) (*models.FolderList, error) {
	return Get[filestation.CreateFolderRequest, models.FolderList](f.client, &filestation.CreateFolderRequest{
		Paths:       paths,
		Names:       names,
		ForceParent: forceParent,
	})
}

// ListShares implements filestation.FileStationApi.
func (f *fileStationClient) ListShares() (*models.ShareList, error) {
	return Get[filestation.ListShareRequest, models.ShareList](f.client, &filestation.ListShareRequest{})
}

// Upload implements filestation.FileStationApi.
func (f *fileStationClient) Upload(path string, file *form.File, createParents bool, overwrite bool) (*filestation.UploadResponse, error) {
	return Post[filestation.UploadRequest, filestation.UploadResponse](f.client, &filestation.UploadRequest{
		Path:          path,
		File:          file,
		CreateParents: createParents,
		Overwrite:     overwrite,
	})
}

func NewFileStationClient(client *APIClient) filestation.FileStationApi {
	return &fileStationClient{client: client}
}
