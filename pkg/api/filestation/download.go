package filestation

import (
	"github.com/synology-community/synology-api/pkg/util/form"
)

type DownloadRequest struct {
	Path string `form:"path" url:"path"`
	Mode string `form:"mode" url:"mode"`
}

type DownloadResponse struct {
	CreateParents bool      `form:"create_parents" url:"create_parents"`
	Overwrite     bool      `form:"overwrite" url:"overwrite"`
	File          form.File `form:"file" kind:"file"`
}
