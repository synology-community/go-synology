package filestation

type MD5StartRequest struct {
	Path string `form:"file_path" url:"file_path"`
}

type MD5StartResponse struct {
	TaskID string `json:"taskid"`
}

type MD5StatusRequest struct {
	TaskID string `url:"taskid,quoted" form:"taskid"`
}

type MD5StatusResponse struct {
	Finished bool   `json:"finished"`
	MD5      string `json:"md5"`
}

type MD5Response struct {
	MD5 string `json:"md5"`
}
