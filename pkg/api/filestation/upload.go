package filestation

import (
	"net/url"

	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type UploadRequest struct {
	Path          string     `form:"path" url:"path"`
	CreateParents bool       `form:"create_parents" url:"create_parents"`
	Overwrite     bool       `form:"overwrite" url:"overwrite"`
	File          *form.File `form:"file" kind:"file"`
}

func (l UploadRequest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}

type UploadResponse struct {
}

var _ api.Request = (*UploadRequest)(nil)

func NewUploadRequest(path string, file *form.File) *UploadRequest {
	return &UploadRequest{
		Path: path,
		File: file,
	}
}

func (r *UploadRequest) WithPath(value string) *UploadRequest {
	r.Path = value
	return r
}

func (r *UploadRequest) WithFile(file *form.File) *UploadRequest {
	r.File = file
	return r
}

func (r *UploadRequest) WithCreateParents(value bool) *UploadRequest {
	r.CreateParents = value
	return r
}

func (r *UploadRequest) WithOverwrite(value bool) *UploadRequest {
	r.Overwrite = value
	return r
}
