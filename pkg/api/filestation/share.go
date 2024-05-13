package filestation

import (
	"net/url"

	"github.com/synology-community/synology-api/pkg/models"
)

type CreateShareRequest struct {
	SortBy     string   `url:"sort_by"`
	FileType   string   `url:"file_type"`
	CheckDir   bool     `url:"check_dir"`
	Additional []string `url:"additional" del:","`
}

type CreateShareResponse struct {
	Offset int `json:"offset"`

	Total int `json:"total"`
}

type ListShareRequest struct {
	SortBy     string   `url:"sort_by"`
	FileType   string   `url:"file_type"`
	CheckDir   bool     `url:"check_dir"`
	Additional []string `url:"additional" del:","`
	GoToPath   string   `url:"goto_path"`
	FolderPath string   `url:"folder_path"`
}

type ListShareResponse struct {
	Offset int `json:"offset"`

	Shares []models.Share `json:"shares"`

	Total int `json:"total"`
}

func (l ListShareRequest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}
