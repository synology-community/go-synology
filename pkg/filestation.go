package client

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/models"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type fileStationClient struct {
	client *APIClient
}

func (f *fileStationClient) Delete(ctx context.Context, paths []string, accurateProgress bool) (*filestation.DeleteStatusResponse, error) {
	// Start Delete the file
	rdel, err := f.client.FileStationAPI().DeleteStart(ctx, paths, true)

	if err != nil {
		return nil, fmt.Errorf("Unable to delete file, got error: %s", err)
	}

	waitUntil := time.Now().Add(60 * time.Second)
	completed := false
	for !completed {
		// Check the status of the delete operation
		rstat, err := f.client.FileStationAPI().DeleteStatus(ctx, rdel.TaskID)
		if err != nil {
			return nil, fmt.Errorf("Unable to delete file, got error: %v", err)
		}

		if rstat.Finished {
			completed = true
			return rstat, nil
		}

		if time.Now().After(waitUntil) {
			return nil, fmt.Errorf("Timeout waiting for file to be deleted")
		}

		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("Unable to delete file, retry count exceeded")
}

func (f *fileStationClient) DeleteStart(ctx context.Context, paths []string, accurateProgress bool) (*filestation.DeleteStartResponse, error) {
	method := filestation.API_METHODS["DeleteStart"]
	return Get[filestation.DeleteStartRequest, filestation.DeleteStartResponse](f.client, ctx, &filestation.DeleteStartRequest{
		Paths:            paths,
		AccurateProgress: accurateProgress,
	}, method)
}

func (f *fileStationClient) DeleteStatus(ctx context.Context, taskID string) (*filestation.DeleteStatusResponse, error) {
	return Get[filestation.DeleteStatusRequest, filestation.DeleteStatusResponse](f.client, ctx, &filestation.DeleteStatusRequest{
		TaskID: taskID,
	}, filestation.API_METHODS["DeleteStatus"])
}

func (f *fileStationClient) MD5Start(ctx context.Context, path string) (*filestation.MD5StartResponse, error) {
	return Get[filestation.MD5StartRequest, filestation.MD5StartResponse](f.client, ctx, &filestation.MD5StartRequest{
		Path: path,
	}, filestation.API_METHODS["MD5Start"])
}

func (f *fileStationClient) MD5Status(ctx context.Context, taskID string) (*filestation.MD5StatusResponse, error) {
	return Get[filestation.MD5StatusRequest, filestation.MD5StatusResponse](f.client, ctx, &filestation.MD5StatusRequest{
		TaskID: taskID,
	}, filestation.API_METHODS["MD5Status"])
}

// Download implements filestation.FileStationApi.
func (f *fileStationClient) Download(ctx context.Context, path string, mode string) (*filestation.DownloadResponse, error) {
	return Get[filestation.DownloadRequest, filestation.DownloadResponse](f.client, ctx, &filestation.DownloadRequest{
		Path: path,
		Mode: mode,
	}, filestation.API_METHODS["Download"])
}

// Rename implements filestation.FileStationApi.
func (f *fileStationClient) Rename(ctx context.Context, path string, name string, newName string) (*models.FileList, error) {
	return Get[filestation.RenameRequest, models.FileList](f.client, ctx, &filestation.RenameRequest{
		Path:    path,
		Name:    name,
		NewName: newName,
	}, filestation.API_METHODS["Rename"])
}

// CreateFolder implements filestation.FileStationApi.
func (f *fileStationClient) CreateFolder(ctx context.Context, paths []string, names []string, forceParent bool) (*models.FolderList, error) {
	return Get[filestation.CreateFolderRequest, models.FolderList](f.client, ctx, &filestation.CreateFolderRequest{
		Paths:       paths,
		Names:       names,
		ForceParent: forceParent,
	}, filestation.API_METHODS["CreateFolder"])
}

// ListShares implements filestation.FileStationApi.
func (f *fileStationClient) ListShares(ctx context.Context) (*models.ShareList, error) {
	return Get[filestation.ListShareRequest, models.ShareList](f.client, ctx, &filestation.ListShareRequest{}, filestation.API_METHODS["ListShares"])
}

func (f *fileStationClient) MD5(ctx context.Context, path string) (*filestation.MD5Response, error) {
	var data filestation.MD5Response
	// Start Delete the file
	rdel, err := f.client.FileStationAPI().MD5Start(ctx, path)

	if err != nil {
		return nil, fmt.Errorf("Unable to delete file, got error: %s", err)
	}

	retry := 0
	completed := false
	for !completed {
		// Check the status of the delete operation
		hstat, err := f.client.FileStationAPI().MD5Status(ctx, rdel.TaskID)
		if err != nil {
			return nil, fmt.Errorf("Unable to get file hash, got error: %s", err)
		}

		if hstat.Finished {
			if hstat.MD5 != "" {
				data.MD5 = hstat.MD5
			}

			completed = true
		}

		if retry > 2 {
			completed = true
			continue
		}
		retry++
		time.Sleep(2 * time.Second)
	}

	if data.MD5 != "" {
		return nil, fmt.Errorf("Unable to get file hash, retry count exceeded")
	} else {
		return &data, nil
	}
}

// Upload implements filestation.FileStationApi.
func (f *fileStationClient) Upload(ctx context.Context, path string, file form.File, createParents bool, overwrite bool) (*filestation.UploadResponse, error) {
	return Post[filestation.UploadRequest, filestation.UploadResponse](f.client, ctx, &filestation.UploadRequest{
		Path:          path,
		File:          file,
		CreateParents: createParents,
		Overwrite:     overwrite,
	}, filestation.API_METHODS["Upload"])
}

func NewFileStationClient(client *APIClient) filestation.FileStationApi {
	return &fileStationClient{client: client}
}
