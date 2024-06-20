package core

type Share struct {
	CompressionRatioTips    bool    `json:"compression_ratio_tips,omitempty"`
	Desc                    string  `json:"desc,omitempty"`
	EnableRecycleBin        bool    `json:"enable_recycle_bin,omitempty"`
	EnableShareCompress     bool    `json:"enable_share_compress,omitempty"`
	EnableShareCow          bool    `json:"enable_share_cow,omitempty"`
	EnableShareTiering      bool    `json:"enable_share_tiering,omitempty"`
	EncAutoMount            bool    `json:"enc_auto_mount,omitempty"`
	Encryption              int     `json:"encryption,omitempty"`
	ForceReadonlyReason     string  `json:"force_readonly_reason,omitempty"`
	Hidden                  bool    `json:"hidden,omitempty"`
	IsAclmode               bool    `json:"is_aclmode,omitempty"`
	IsApplyingSettings      bool    `json:"is_applying_settings,omitempty"`
	IsBlockSnapAction       bool    `json:"is_block_snap_action,omitempty"`
	IsC2Share               bool    `json:"is_c2_share,omitempty"`
	IsClusterShare          bool    `json:"is_cluster_share,omitempty"`
	IsColdStorageShare      bool    `json:"is_cold_storage_share,omitempty"`
	IsExfatShare            bool    `json:"is_exfat_share,omitempty"`
	IsForceReadonly         bool    `json:"is_force_readonly,omitempty"`
	IsMissingShare          bool    `json:"is_missing_share,omitempty"`
	IsOfflineShare          bool    `json:"is_offline_share,omitempty"`
	IsShareMoving           bool    `json:"is_share_moving,omitempty"`
	IsSupportACL            bool    `json:"is_support_acl,omitempty"`
	IsSyncShare             bool    `json:"is_sync_share,omitempty"`
	IsUsbShare              bool    `json:"is_usb_share,omitempty"`
	Name                    string  `json:"name,omitempty"`
	QuotaValue              int     `json:"quota_value,omitempty"`
	RecycleBinAdminOnly     bool    `json:"recycle_bin_admin_only,omitempty"`
	ShareQuotaLogicalSize   int     `json:"share_quota_logical_size,omitempty"`
	ShareQuotaStatus        string  `json:"share_quota_status,omitempty"`
	ShareQuotaUsed          float64 `json:"share_quota_used,omitempty"`
	SupportAction           int     `json:"support_action,omitempty"`
	SupportCompressionRatio bool    `json:"support_compression_ratio,omitempty"`
	SupportSnapshot         bool    `json:"support_snapshot,omitempty"`
	TaskID                  string  `json:"task_id,omitempty"`
	UnitePermission         bool    `json:"unite_permission,omitempty"`
	UUID                    string  `json:"uuid,omitempty"`
	VolPath                 string  `json:"vol_path,omitempty"`
	WormDefLockDuration     int     `json:"worm_def_lock_duration,omitempty"`
	WormDefLockMode         string  `json:"worm_def_lock_mode,omitempty"`
	WormLockWaitTime        int     `json:"worm_lock_wait_time,omitempty"`
	WormSubvolMode          string  `json:"worm_subvol_mode,omitempty"`
}

type ShareList struct {
	Shares []Share `json:"shares,omitempty"`
	Total  *int    `json:"total,omitempty"`
}

type ShareListRequest struct {
	ShareType  string   `url:"share_type,omitempty"`
	Additional []string `url:"additional,json"`
}

type ShareListResponse ShareList

type ShareGetRequest struct {
	Name       string   `url:"name"`
	Additional []string `url:"additional,json"`
}

type ShareGetResponse Share

type ShareInfo struct {
	Name                string `json:"name,omitempty"`
	VolPath             string `json:"vol_path,omitempty"`
	Desc                string `json:"desc,omitempty"`
	EnableShareCow      bool   `json:"enable_share_cow,omitempty"`
	EnableShareCompress bool   `json:"enable_share_compress,omitempty"`
	NameOrg             string `json:"name_org,omitempty"`
}

type ShareCreateRequest struct {
	Name      string    `url:"name"`
	ShareInfo ShareInfo `url:"shareinfo,json"`
}

type ShareDeleteRequest struct {
	Name string `url:"name"`
}
