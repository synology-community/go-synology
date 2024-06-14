package docker

import (
	"net/url"

	"github.com/synology-community/go-synology/pkg/util"
)

type ImagePullStartRequest struct {
	Repository string `form:"repository" url:"repository,quoted"`
	Tag        string `form:"tag" url:"tag,quoted"`
}

type ImagePullStartResponse struct {
	TaskID string `json:"task_id"`
}

type ImagePullStatusRequest struct {
	TaskID string `form:"task_id" url:"task_id,quoted"`
}

type ImagePullStatusResponse struct {
	Description string `json:"description"`
	Downloaded  int64  `json:"downloaded"`
	Finished    bool   `json:"finished"`
	Repository  string `json:"repository"`
	Tag         string `json:"tag"`
}

type Image struct {
	Repository string   `json:"repository"`
	Tags       []string `json:"tags"`
}

type ImageList []Image

type ImageDeleteRequest struct {
	Images ImageList `json:"images" url:"images"`
}

type ImageDeleteResponse struct {
	ImageObjects map[string]map[string]struct {
		Error int64 `json:"error"`
	} `json:"image_objects"`
}

func (i ImageList) EncodeValues(k string, v *url.Values) error {
	return util.EncodeValues(&i, k, v)
}
