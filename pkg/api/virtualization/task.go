package virtualization

type TaskInfo struct {
	Progress  int    `json:"progress"`
	Status    string `json:"status"`
	AutoClean bool   `json:"auto_clean_task,omitempty"`
	ImageID   string `json:"image_id,omitempty"`
	GuestID   string `json:"guest_id,omitempty"`
}

type Task struct {
	Finished bool     `json:"finish"`
	TaskInfo TaskInfo `json:"task_info"`
}

type TaskRef struct {
	TaskID string `url:"task_id" json:"task_id"`
}
