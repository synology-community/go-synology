package core

import "github.com/synology-community/synology-api/pkg/models"

type PackageListRequest struct {
	IgnoreHidden bool     `url:"ignore_hidden"`
	Additional   []string `url:"additional"`
}

type PackageListResponse struct {
	Packages []struct {
		Additional struct {
			Autoupdate          bool `json:"autoupdate"`
			AutoupdateImportant bool `json:"autoupdate_important"`
			AvailableOperation  struct {
			} `json:"available_operation"`
			Beta              bool   `json:"beta"`
			CtlUninstall      bool   `json:"ctl_uninstall"`
			DependentPackages any    `json:"dependent_packages"`
			Description       string `json:"description"`
			DescriptionEnu    string `json:"description_enu"`
			Distributor       string `json:"distributor"`
			DistributorURL    string `json:"distributor_url"`
			DsmAppLaunchName  string `json:"dsm_app_launch_name"`
			DsmAppPage        string `json:"dsm_app_page"`
			DsmApps           string `json:"dsm_apps"`
			InstallType       string `json:"install_type"`
			InstalledInfo     struct {
				IsBrick  bool   `json:"is_brick"`
				IsBroken bool   `json:"is_broken"`
				Path     string `json:"path"`
			} `json:"installed_info"`
			IsUninstallPages  bool   `json:"is_uninstall_pages"`
			Maintainer        string `json:"maintainer"`
			MaintainerURL     string `json:"maintainer_url"`
			ReportBetaURL     string `json:"report_beta_url"`
			SilentUpgrade     bool   `json:"silent_upgrade"`
			Startable         bool   `json:"startable"`
			Status            string `json:"status"`
			StatusCode        int    `json:"status_code"`
			StatusDescription string `json:"status_description"`
			StatusOrigin      string `json:"status_origin"`
			SupportCenter     bool   `json:"support_center"`
			SupportURL        string `json:"support_url"`
			UpdatedAt         string `json:"updated_at"`
			URL               []any  `json:"url"`
		} `json:"additional"`
		ID        string `json:"id"`
		Name      string `json:"name"`
		Timestamp int64  `json:"timestamp"`
		Version   string `json:"version"`
	} `json:"packages"`
	Total int `json:"total"`
}

type PackageGetRequest struct {
	ID models.JsonString `url:"id"`
}

type PackageGetResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Version   string `json:"version,omitempty"`
}

type PackageServerListRequest struct {
	ForceReload bool `url:"blforcereload"`
	LoadOthers  bool `url:"blloadothers"`
}

type Package struct {
	ID                  string   `json:"id"`
	Link                string   `json:"link"`
	Beta                bool     `json:"beta"`
	Breakpkgs           any      `json:"breakpkgs"`
	Changelog           string   `json:"changelog"`
	Conflictpkgs        any      `json:"conflictpkgs"`
	Deppkgs             any      `json:"deppkgs"`
	Depsers             string   `json:"depsers"`
	Desc                string   `json:"desc"`
	Distributor         string   `json:"distributor"`
	DistributorURL      string   `json:"distributor_url"`
	Dname               string   `json:"dname"`
	DownloadCount       int      `json:"download_count"`
	Maintainer          string   `json:"maintainer"`
	MaintainerURL       string   `json:"maintainer_url"`
	Package             string   `json:"package"`
	Qinst               bool     `json:"qinst"`
	Qstart              bool     `json:"qstart"`
	Qupgrade            bool     `json:"qupgrade"`
	RecentDownloadCount int      `json:"recent_download_count,omitempty"`
	Replaceforcepkgs    any      `json:"replaceforcepkgs,omitempty"`
	Replacepkgs         any      `json:"replacepkgs,omitempty"`
	Source              string   `json:"source"`
	Thumbnail           []string `json:"thumbnail,omitempty"`
	Version             string   `json:"version"`
	MD5                 string   `json:"md5,omitempty"`
	Size                int64    `json:"size,omitempty"`
}

type PackageServerListResponse struct {
	BetaPackages []any     `json:"beta_packages"`
	Packages     []Package `json:"packages"`
}

type PackageInstallStatusRequest struct {
	TaskID string `url:"task_id"`
}

type PackageInstallStatusResponse struct {
	Beta       bool   `json:"beta,omitempty"`
	Blqinst    bool   `json:"blqinst,omitempty"`
	ID         string `json:"id,omitempty"`
	Installing bool   `json:"installing,omitempty"`
	Name       string `json:"name,omitempty"`
	Pid        int    `json:"pid,omitempty"`
	RemoteLink string `json:"remote_link,omitempty"`
	Size       string `json:"size,omitempty"`
	Status     string `json:"status,omitempty"`
	Success    bool   `json:"success,omitempty"`
	Taskid     string `json:"taskid,omitempty"`
	TmpFolder  string `json:"tmp_folder,omitempty"`
	Version    string `json:"version,omitempty"`
	Finished   bool   `json:"finished"`
}

type PackageInstallRequest struct {
	Name              string            `url:"name"`
	URL               string            `url:"url,omitempty"`
	Type              int64             `url:"type,omitempty"`
	BigInstall        bool              `url:"blqinst,omitempty"`
	Checksum          models.JsonString `url:"checksum,omitempty"`
	FileSize          int64             `url:"filesize,omitempty"`
	ExtraValues       models.JsonString `url:"extra_values,omitempty"`
	CheckCodesign     bool              `url:"check_codesign,omitempty"`
	Force             bool              `url:"force,omitempty"`
	InstallRunPackage bool              `url:"installrunpackage,omitempty"`
	Path              string            `url:"path,omitempty"`
}

type PackageInstallResponse struct {
	TaskID   string  `json:"taskid"`
	Progress float64 `json:"progress,omitempty"`
}
