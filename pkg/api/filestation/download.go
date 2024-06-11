package filestation

type DownloadRequest struct {
	Path string `form:"path" url:"path"`
	Mode string `form:"mode" url:"mode"`
}
