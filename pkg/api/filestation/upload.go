package filestation

import (
	"net/url"

	"github.com/synology-community/go-synology/internal/util/form"
	"github.com/synology-community/go-synology/pkg/api"
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
