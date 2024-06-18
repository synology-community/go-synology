package models

import (
	"fmt"
	"strconv"
	"time"
)

type Pagination struct {
	Total  int `json:"total,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type FolderList struct {
	Folders []File `json:"folders"`
}

type FileListRequest struct {
	FolderPath string   `url:"folder_path"`
	Additional []string `url:"additional,json"`
	FileType   string   `url:"filetype,omitempty"`
}

type FileList struct {
	Pagination
	Files []File `json:"files"`
}

type ShareList struct {
	Pagination
	Shares []Share `json:"shares"`
}

type Share struct {
	Name       string               `json:"name"`
	Path       string               `json:"path"`
	IsDir      bool                 `json:"isdir"`
	Additional AdditionalAttributes `json:"additional,omitempty"`
}

type File struct {
	Path       string               `json:"path"`
	Name       string               `json:"name"`
	IsDir      bool                 `json:"isdir"`
	Additional AdditionalAttributes `json:"additional,omitempty"`
}

type AdditionalAttributes struct {
	Indexed        bool   `json:"indexed,omitempty"`
	IsHybridShare  bool   `json:"is_hybrid_share,omitempty"`
	IsWormShare    bool   `json:"is_worm_share,omitempty"`
	MountPointType string `json:"mount_point_type,omitempty"`
	Owner          struct {
		Group   string `json:"group,omitempty"`
		GroupID int    `json:"gid,omitempty"`
		User    string `json:"user,omitempty"`
		UserID  int    `json:"uid,omitempty"`
	} `json:"owner,omitempty"`
	Perm struct {
		ACL struct {
			Append bool `json:"append,omitempty"`
			Del    bool `json:"del,omitempty"`
			Exec   bool `json:"exec,omitempty"`
			Read   bool `json:"read,omitempty"`
			Write  bool `json:"write,omitempty"`
		} `json:"acl,omitempty"`
		ACLEnable bool `json:"acl_enable,omitempty"`
		AdvRight  struct {
			DisableDownload bool `json:"disable_download,omitempty"`
			DisableList     bool `json:"disable_list,omitempty"`
			DisableModify   bool `json:"disable_modify,omitempty"`
		} `json:"adv_right,omitempty"`
		IsACLMode       bool   `json:"is_acl_mode"`
		IsShareReadonly bool   `json:"is_share_readonly"`
		Posix           int    `json:"posix"`
		ShareRight      string `json:"share_right"`
	} `json:"perm,omitempty"`
	RealPath  string `json:"real_path,omitempty"`
	SyncShare bool   `json:"sync_share,omitempty"`
	Time      struct {
		Atime  Time `json:"atime"`
		Crtime Time `json:"crtime"`
		Ctime  Time `json:"ctime"`
		Mtime  Time `json:"mtime"`
	} `json:"time,omitempty"`
	VolumeStatus struct {
		Freespace  int64 `json:"freespace,omitempty"`
		Readonly   bool  `json:"readonly,omitempty"`
		Totalspace int64 `json:"totalspace,omitempty"`
	} `json:"volume_status,omitempty"`
	WormState int `json:"worm_state,omitempty"`
}

type Time struct {
	time.Time
}

func (ms Time) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`%d`, ms.Unix())
	return []byte(str), nil
}

func (ms Time) RFC3339() string {
	return ms.Format(time.RFC3339)
}

func (ms *Time) UnmarshalJSON(text []byte) error {
	i, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	*ms = Time{Time: time.Unix(i, 0)}
	return nil
}
