package filestation

import (
	"net/url"

	"github.com/synology-community/synology-api/package/models"
)

type CreateShareRequest struct {
	SortBy     string   `query:"sort_by"`
	FileType   string   `query:"file_type"`
	CheckDir   bool     `query:"check_dir"`
	Additional []string `query:"additional" del:","`
}

type CreateShareResponse struct {
	Offset int `json:"offset"`

	Total int `json:"total"`
}

type ListShareRequest struct {
	SortBy     string   `query:"sort_by"`
	FileType   string   `query:"file_type"`
	CheckDir   bool     `query:"check_dir"`
	Additional []string `query:"additional" del:","`
	GoToPath   string   `query:"goto_path"`
	FolderPath string   `query:"folder_path"`
}

type ListShareResponse struct {
	Offset int `json:"offset"`

	Shares []models.Share `json:"shares"`

	Total int `json:"total"`
}

func (l ListShareRequest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}

// func NewListShareRequest(sortBy string, fileType string, checkDir bool, additional []string, goToPath string, folderPath string) *ListShareRequest {
//
// 	if additional == nil {
// 		additional = []string{"real_path", "owner", "time", "perm", "mount_point_type", "sync_share", "volume_status", "indexed", "hybrid_share", "worm_share"}
// 	}
// 	if sortBy == "" {
// 		sortBy = "name"
// 	}
// 	return &ListShareRequest{
// 		SortBy:     sortBy,
// 		FileType:   fileType,
// 		CheckDir:   checkDir,
// 		Additional: additional,
// 		GoToPath:   goToPath,
// 		FolderPath: folderPath,
// 	}
// }
