package filestation

import "net/url"

type CreateFolderRequest struct {
	Paths       []string `form:"folder_path" url:"folder_path"`
	Names       []string `form:"name" url:"name"`
	ForceParent bool     `form:"force_parent" url:"force_parent"`
}

func (l CreateFolderRequest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}

type CreateFolderResponse struct {
	Folders []struct {
		Path  string
		Name  string
		IsDir bool
	}
}
