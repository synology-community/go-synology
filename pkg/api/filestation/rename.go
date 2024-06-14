package filestation

type RenameRequest struct {
	Name    string `url:"name"`
	Path    string `url:"path"`
	NewName string `url:"new_name"`
}
