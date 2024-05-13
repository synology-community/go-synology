package filestation

import "net/url"

type RenameRequest struct {
	Name    string `url:"name"`
	Path    string `url:"path"`
	NewName string `url:"new_name"`
}

func (l RenameRequest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}
