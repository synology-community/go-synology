package filestation

type DeleteStartRequest struct {
	Paths            []string `form:"path" url:"path"`
	AccurateProgress bool     `form:"accurate_progress" url:"accurate_progress"`
}

type DeleteStartResponse struct {
	TaskID string `json:"taskid"`
}

type DeleteStatusRequest struct {
	TaskID string `form:"taskid" url:"taskid"`
}

type DeleteStatusResponse struct {
	Finished       bool   `json:"finished"`
	FoundDirNum    int    `json:"found_dir_num"`
	FoundFileNum   int    `json:"found_file_num"`
	FoundFileSize  int    `json:"found_file_size"`
	HasDir         bool   `json:"has_dir"`
	Path           string `json:"path"`
	ProcessedNum   int    `json:"processed_num"`
	ProcessingPath string `json:"processing_path"`
	Progress       int    `json:"progress"`
	Total          int    `json:"total"`
}
