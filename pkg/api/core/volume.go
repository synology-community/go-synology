package core

type VolumeListRequest struct {
	Limit    int    `url:"limit"`
	Offset   int    `url:"offset"`
	Location string `url:"location,quoted"`
}

type Volume struct {
	AtimeChecked              bool   `json:"atime_checked,omitempty"`
	AtimeOpt                  string `json:"atime_opt,omitempty"`
	Container                 string `json:"container,omitempty"`
	Crashed                   bool   `json:"crashed,omitempty"`
	Deduped                   bool   `json:"deduped,omitempty"`
	Description               string `json:"description,omitempty"`
	DisplayName               string `json:"display_name,omitempty"`
	FsType                    string `json:"fs_type,omitempty"`
	IsEncrypted               bool   `json:"is_encrypted,omitempty"`
	Location                  string `json:"location,omitempty"`
	PoolPath                  string `json:"pool_path,omitempty"`
	RaidType                  string `json:"raid_type,omitempty"`
	Readonly                  bool   `json:"readonly,omitempty"`
	SingleVolume              bool   `json:"single_volume,omitempty"`
	SizeFreeByte              string `json:"size_free_byte,omitempty"`
	SizeTotalByte             string `json:"size_total_byte,omitempty"`
	Status                    string `json:"status,omitempty"`
	VolumeAttribute           string `json:"volume_attribute,omitempty"`
	VolumeID                  int    `json:"volume_id,omitempty"`
	VolumePath                string `json:"volume_path,omitempty"`
	VolumeQuotaStatus         string `json:"volume_quota_status,omitempty"`
	VolumeQuotaUpdateProgress int    `json:"volume_quota_update_progress,omitempty"`
}

type VolumeListResponse struct {
	Offset  int      `json:"offset,omitempty"`
	Total   int      `json:"total,omitempty"`
	Volumes []Volume `json:"volumes,omitempty"`
}
