package filestation

type DeleteStartRequest struct {
	Paths            []string `form:"path" url:"path"`
	AccurateProgress bool     `form:"accurate_progress" url:"accurate_progress"`
}

type DeleteStartResponse struct {
	TaskID string `mapstructure:"taskid" json:"taskid"`
}

type DeleteStatusRequest struct {
	TaskID string `form:"taskid" url:"taskid"`
}

type DeleteStatusResponse struct {
	Finished       bool   `mapstructure:"finished" json:"finished"`
	FoundDirNum    int    `mapstructure:"found_dir_num" json:"found_dir_num"`
	FoundFileNum   int    `mapstructure:"found_file_num" json:"found_file_num"`
	FoundFileSize  int    `mapstructure:"found_file_size" json:"found_file_size"`
	HasDir         bool   `mapstructure:"has_dir" json:"has_dir"`
	Path           string `mapstructure:"path" json:"path"`
	ProcessedNum   int    `mapstructure:"processed_num" json:"processed_num"`
	ProcessingPath string `mapstructure:"processing_path" json:"processing_path"`
	Progress       int    `mapstructure:"progress" json:"progress"`
	Total          int    `mapstructure:"total" json:"total"`
}
