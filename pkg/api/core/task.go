package core

type CreateTaskRequest struct {
	SortBy     string   `form:"sort_by" url:"sort_by"`
	FileType   string   `form:"file_type" url:"file_type"`
	CheckDir   bool     `form:"check_dir" url:"check_dir"`
	Additional []string `form:"additional" url:"additional"`
}

type CreateTaskResponse struct {
	Offset int `json:"offset"`

	Total int `json:"total"`
}

type ListTaskRequest struct {
	SortBy     string   `form:"sort_by" url:"sort_by"`
	FileType   string   `form:"file_type" url:"file_type"`
	CheckDir   bool     `form:"check_dir" url:"check_dir"`
	Additional []string `form:"additional" url:"additional"`
	GoToPath   string   `form:"goto_path" url:"goto_path"`
	FolderPath string   `form:"folder_path" url:"folder_path"`
}

type ListTaskResponse struct {
	Offset int `json:"offset"`

	Tasks []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		IsDir      bool   `json:"isdir"`
		Additional struct {
			Indexed        bool   `json:"indexed"`
			IsHybridTask   bool   `json:"is_hybrid_share"`
			IsWormTask     bool   `json:"is_worm_share"`
			MountPointType string `json:"mount_point_type"`
			Owner          struct {
				Group   string `json:"group"`
				GroupID int    `json:"gid"`
				User    string `json:"user"`
				UserID  int    `json:"uid"`
			} `json:"owner"`
			Perm struct {
				ACL struct {
					Append bool `json:"append"`
					Del    bool `json:"del"`
					Exec   bool `json:"exec"`
					Read   bool `json:"read"`
					Write  bool `json:"write"`
				} `json:"acl"`
				ACLEnable bool `json:"acl_enable"`
				AdvRight  struct {
					DisableDownload bool `json:"disable_download"`
					DisableList     bool `json:"disable_list"`
					DisableModify   bool `json:"disable_modify"`
				} `json:"adv_right"`
				IsACLMode      bool   `json:"is_acl_mode"`
				IsTaskReadonly bool   `json:"is_share_readonly"`
				Posix          int    `json:"posix"`
				TaskRight      string `json:"share_right"`
			} `json:"perm"`
			RealPath string `json:"real_path"`
			SyncTask bool   `json:"sync_share"`
			Time     struct {
				Atime  int `json:"atime"`
				Crtime int `json:"crtime"`
				Ctime  int `json:"ctime"`
				Mtime  int `json:"mtime"`
			} `json:"time"`
			VolumeStatus struct {
				Freespace  int64 `json:"freespace"`
				Readonly   bool  `json:"readonly"`
				Totalspace int64 `json:"totalspace"`
			} `json:"volume_status"`
			WormState int `json:"worm_state"`
		} `json:"additional"`
	} `json:"shares"`

	Total int `json:"total"`
}
