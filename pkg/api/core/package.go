package core

import "github.com/synology-community/go-synology/pkg/models"

type PackageListRequest struct {
	IgnoreHidden bool             `url:"ignore_hidden"`
	Additional   models.JsonArray `url:"additional"`
}

type InstalledPackage struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	Version    string `json:"version,omitempty"`
	Additional struct {
		Autoupdate          bool `json:"autoupdate,omitempty"`
		AutoupdateImportant bool `json:"autoupdate_important,omitempty"`
		AvailableOperation  struct {
		} `json:"available_operation,omitempty"`
		Beta              bool   `json:"beta,omitempty"`
		CtlUninstall      bool   `json:"ctl_uninstall,omitempty"`
		DependentPackages any    `json:"dependent_packages,omitempty"`
		Description       string `json:"description,omitempty"`
		DescriptionEnu    string `json:"description_enu,omitempty"`
		Distributor       string `json:"distributor,omitempty"`
		DistributorURL    string `json:"distributor_url,omitempty"`
		DsmAppLaunchName  string `json:"dsm_app_launch_name,omitempty"`
		DsmAppPage        string `json:"dsm_app_page,omitempty"`
		DsmApps           string `json:"dsm_apps,omitempty"`
		InstallType       string `json:"install_type,omitempty"`
		InstalledInfo     struct {
			IsBrick  bool   `json:"is_brick,omitempty"`
			IsBroken bool   `json:"is_broken,omitempty"`
			Path     string `json:"path,omitempty"`
		} `json:"installed_info,omitempty"`
		IsUninstallPages  bool   `json:"is_uninstall_pages,omitempty"`
		Maintainer        string `json:"maintainer,omitempty"`
		MaintainerURL     string `json:"maintainer_url,omitempty"`
		ReportBetaURL     string `json:"report_beta_url,omitempty"`
		SilentUpgrade     bool   `json:"silent_upgrade,omitempty"`
		Startable         bool   `json:"startable,omitempty"`
		Status            string `json:"status,omitempty"`
		StatusCode        int    `json:"status_code,omitempty"`
		StatusDescription string `json:"status_description,omitempty"`
		StatusOrigin      string `json:"status_origin,omitempty"`
		SupportCenter     bool   `json:"support_center,omitempty"`
		SupportURL        string `json:"support_url,omitempty"`
		UpdatedAt         string `json:"updated_at,omitempty"`
		URL               []any  `json:"url,omitempty"`
	} `json:"additional,omitempty"`
}

type PackageListResponse struct {
	Packages []InstalledPackage `json:"packages"`
	Total    int                `json:"total"`
}

type PackageGetRequest struct {
	ID         models.JsonString `url:"id"`
	Additional models.JsonArray  `url:"additional,omitempty"`
}

type PackageGetResponse InstalledPackage

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

type PackageInstallCheckRequest struct {
	ID                   string `url:"id"`
	InstallType          string `url:"install_type"`
	InstallOnColdStorage bool   `url:"install_on_cold_storage"`
	BreakPkgs            string `url:"breakpkgs"`
	BlCheckDep           bool   `url:"blcheckdep"`
	ReplacePkgs          string `url:"replacepkgs"`
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
	Checksum          models.JsonString `url:"checksum"`
	FileSize          int64             `url:"filesize,omitempty"`
	ExtraValues       models.JsonString `url:"extra_values,omitempty"`
	CheckCodesign     bool              `url:"check_codesign,omitempty"`
	Force             bool              `url:"force,omitempty"`
	InstallRunPackage bool              `url:"installrunpackage,omitempty"`
	Path              models.JsonString `url:"path,omitempty"`
	Operation         models.JsonString `url:"operation,omitempty"`
	VolumePath        string            `url:"volume_path,omitempty"`
}

type PackageInstallResponse struct {
	TaskID   string  `json:"taskid"`
	Progress float64 `json:"progress,omitempty"`
}

type PackageFindRequest struct {
	Name string `url:"name"`
}

type PackageUninstallRequest struct {
	ID      string `url:"id"`
	DSMApps string `url:"dsm_apps,omitempty"`
}

type PackageUninstallResponse struct {
	Message       string   `json:"message,omitempty"`
	WorkerMessage []string `json:"worker_message,omitempty"`
}
