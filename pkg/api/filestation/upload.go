package filestation

import (
	"net/url"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/util/form"
)

type UploadRequest struct {
	Path          string    `form:"path" url:"path"`
	CreateParents bool      `form:"create_parents" url:"create_parents"`
	Overwrite     bool      `form:"overwrite" url:"overwrite"`
	File          form.File `form:"file" kind:"file"`
}

func (l UploadRequest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}

type UploadResponse struct {
}

var _ api.Request = (*UploadRequest)(nil)
