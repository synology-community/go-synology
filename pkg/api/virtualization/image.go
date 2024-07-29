package virtualization

import (
	"fmt"
	"net/url"

	"github.com/synology-community/go-synology/pkg/util/form"
)

type ImageType string

func (s ImageType) EncodeValues(k string, v *url.Values) error {
	if s == "disk" || s == "iso" || s == "vdsm" {
		v.Set(k, string(s))
	} else {
		return fmt.Errorf("invalid image type: %s. Must be one of disk, iso, or vdsm", s)
	}
	return nil
}

type Image struct {
	ID          string    `url:"image_id,omitempty" json:"image_id"`
	Name        string    `url:"image_name,omitempty" json:"image_name"`
	FilePath    string    `url:"ds_file_path,omitempty" json:"-"`
	Description string    `url:"description,omitempty" json:"description,omitempty"`
	Storages    Storages  `url:"storages,omitempty" json:"storages"`
	Type        ImageType `url:"type,omitempty" json:"type"`
	AutoClean   bool      `url:"auto_clean_task,omitempty" json:"-"`
}

type ImageList struct {
	Images []Image `json:"images"`
}

type UploadAndCreateRequest struct {
	Name       string    `url:"name" form:"-"`
	DsFilePath string    `url:"ds_file_path" form:"-"`
	ImageRepos []string  `url:"image_repos,json,quoted" form:"-"`
	Type       string    `url:"type" form:"-"`
	GetPatchBy string    `url:"get_patch_by" form:"-"`
	File       form.File `url:"-" form:"file" kind:"file"`
}
