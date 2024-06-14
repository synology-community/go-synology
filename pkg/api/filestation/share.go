package filestation

import (
	"github.com/synology-community/go-synology/pkg/models"
)

type CreateShareRequest struct {
	SortBy     string   `url:"sort_by"`
	FileType   string   `url:"file_type"`
	CheckDir   bool     `url:"check_dir"`
	Additional []string `url:"additional,json"`
}

type CreateShareResponse struct {
	Offset int `json:"offset"`

	Total int `json:"total"`
}

type ListShareRequest struct {
	SortBy     string   `url:"sort_by"`
	FileType   string   `url:"file_type"`
	CheckDir   bool     `url:"check_dir"`
	Additional []string `url:"additional,json"`
	GoToPath   string   `url:"goto_path"`
	FolderPath string   `url:"folder_path"`
}

type ListShareResponse struct {
	Offset int `json:"offset"`

	Shares []models.Share `json:"shares"`

	Total int `json:"total"`
}
