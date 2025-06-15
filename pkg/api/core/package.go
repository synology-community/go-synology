package core

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/synology-community/go-synology/pkg/util"
	"github.com/synology-community/go-synology/pkg/util/form"
)

type PackageListRequest struct {
	IgnoreHidden bool     `url:"ignore_hidden"`
	Additional   []string `url:"additional,json"`
}

type InstalledPackage struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	Version    string `json:"version,omitempty"`
	Additional struct {
		Autoupdate          bool     `json:"autoupdate,omitempty"`
		AutoupdateImportant bool     `json:"autoupdate_important,omitempty"`
		AvailableOperation  struct{} `json:"available_operation,omitempty"`
		Beta                bool     `json:"beta,omitempty"`
		CtlUninstall        bool     `json:"ctl_uninstall,omitempty"`
		DependentPackages   any      `json:"dependent_packages,omitempty"`
		Description         string   `json:"description,omitempty"`
		DescriptionEnu      string   `json:"description_enu,omitempty"`
		Distributor         string   `json:"distributor,omitempty"`
		DistributorURL      string   `json:"distributor_url,omitempty"`
		DsmAppLaunchName    string   `json:"dsm_app_launch_name,omitempty"`
		DsmAppPage          string   `json:"dsm_app_page,omitempty"`
		DsmApps             string   `json:"dsm_apps,omitempty"`
		InstallType         string   `json:"install_type,omitempty"`
		InstalledInfo       struct {
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
	ID         string   `url:"id"`
	Additional []string `url:"additional,omitempty,json"`
}

type PackageGetResponse InstalledPackage

type PackageServerListRequest struct {
	ForceReload bool `url:"blforcereload"`
	LoadOthers  bool `url:"blloadothers"`
}

type CInt int64

type Package struct {
	ID                  string   `json:"id,omitempty"`
	Link                string   `json:"link,omitempty"`
	Beta                bool     `json:"beta,omitempty"`
	Breakpkgs           any      `json:"breakpkgs,omitempty"`
	Changelog           string   `json:"changelog,omitempty"`
	Conflictpkgs        any      `json:"conflictpkgs,omitempty"`
	Deppkgs             any      `json:"deppkgs,omitempty"`
	Depsers             string   `json:"depsers,omitempty"`
	Desc                string   `json:"desc,omitempty"`
	Distributor         string   `json:"distributor,omitempty"`
	DistributorURL      string   `json:"distributor_url,omitempty"`
	Dname               string   `json:"dname,omitempty"`
	DownloadCount       CInt     `json:"download_count,omitempty"`
	Maintainer          string   `json:"maintainer,omitempty"`
	MaintainerURL       string   `json:"maintainer_url,omitempty"`
	Package             string   `json:"package,omitempty"`
	Qinst               bool     `json:"qinst,omitempty"`
	Qstart              bool     `json:"qstart,omitempty"`
	Qupgrade            bool     `json:"qupgrade,omitempty"`
	RecentDownloadCount int64    `json:"recent_download_count,omitempty"`
	Replaceforcepkgs    any      `json:"replaceforcepkgs,omitempty"`
	Replacepkgs         any      `json:"replacepkgs,omitempty"`
	Source              string   `json:"source,omitempty"`
	Thumbnail           []string `json:"thumbnail,omitempty"`
	Version             string   `json:"version,omitempty"`
	MD5                 string   `json:"md5,omitempty"`
	Size                int64    `json:"size,omitempty"`
}

func (t *CInt) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	s := string(data)
	if s == "null" || s == `""` {
		return nil
	}
	if strings.Contains(s, `"`) {
		s = strings.Trim(s, `"`)
	}
	// Fractional seconds are handled implicitly by Parse.
	i, err := strconv.ParseInt(s, 10, 64)
	*t = CInt(i)
	return err
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

type PackageInstallCheckResponse struct{}

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
	Name              string      `url:"name"`
	URL               string      `url:"url,omitempty"`
	Type              int64       `url:"type,omitempty"`
	BigInstall        bool        `url:"blqinst,omitempty"`
	Checksum          []string    `url:"checksum,quoted"`
	FileSize          int64       `url:"filesize,omitempty"`
	ExtraValues       ExtraValues `url:"extra_values,omitempty"`
	CheckCodesign     bool        `url:"check_codesign"`
	Force             bool        `url:"force,omitempty"`
	InstallRunPackage bool        `url:"installrunpackage"`
	Path              string      `url:"path,omitempty"`
	Operation         []string    `url:"operation,omitempty,quoted"`
	VolumePath        string      `url:"volume_path,omitempty"`
}

type PackageInstallResponse struct {
	TaskID      string  `json:"taskid"`
	Progress    float64 `json:"progress,omitempty"`
	PackageName string  `json:"packageName,omitempty"`
}

type PackageInstallDeleteRequest struct {
	Path string `url:"path"`
}

type PackageFindRequest struct {
	Name string `url:"name"`
}

type UninstallExtra struct {
	KeepData   bool `json:"wizard_keep_data"`
	DeleteData bool `json:"wizard_delete_data"`
}

func (s UninstallExtra) EncodeValues(k string, v *url.Values) error {
	return util.EncodeValues(s, k, v)
}

type PackageUninstallRequest struct {
	ID          string         `url:"id"`
	DSMApps     string         `url:"dsm_apps,omitempty"`
	ExtraValues UninstallExtra `url:"extra_values,omitempty"`
}

type PackageUninstallResponse struct {
	Message       string   `json:"message,omitempty"`
	WorkerMessage []string `json:"worker_message,omitempty"`
}

type PackageFeedItem struct {
	Feed string `json:"feed"`
	Name string `json:"name"`
}

func (s PackageFeedItem) EncodeValues(k string, v *url.Values) error {
	return util.EncodeValuesWrap(s, k, v)
}

type PackageFeeds []string

func (s PackageFeeds) EncodeValues(k string, v *url.Values) error {
	return util.EncodeValuesWrap(s, k, v)
}

type PackageFeedListResponse struct {
	Items []PackageFeedItem `json:"items"`
	Total int64             `json:"total"`
}

type PackageFeedAddRequest struct {
	List PackageFeedItem `url:"list"`
}

type PackageFeedDeleteRequest struct {
	List PackageFeeds `url:"list"`
}

type PackageSettingGetRequest struct {
	Option []string `url:"option,quoted"`
}

type ExtraValues map[string]string

func (s ExtraValues) EncodeValues(k string, v *url.Values) error {
	if len(s) == 0 {
		v.Set(k, `"{}"`)
		return nil
	}

	conf := make(map[string]string)
	for k, v := range s {
		conf[k] = v
		// if !strings.HasPrefix(k, "pkgwizard_") && !strings.HasPrefix(k, "wizard_") {
		// 	conf["wizard_"+k] = v
		// } else {
		// 	conf[k] = v
		// }
	}

	encoded, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	v.Set(k, fmt.Sprintf(`"%s"`, encoded))
	return nil
}

type PackageInstallCompoundRequest struct {
	Name        string            `url:"name"`
	File        string            `url:"file"`
	URL         string            `url:"url"`
	Size        int64             `url:"size"`
	Run         bool              `url:"run"`
	ExtraValues map[string]string `url:"extra_values"`
}

type PackageSettingGetResponse struct {
	Autoupdateall       bool   `json:"autoupdateall,omitempty"`
	Autoupdateimportant bool   `json:"autoupdateimportant,omitempty"`
	DefaultVol          string `json:"default_vol,omitempty"`
	EnableAutoupdate    bool   `json:"enable_autoupdate,omitempty"`
	EnableDsm           bool   `json:"enable_dsm,omitempty"`
	EnableEmail         bool   `json:"enable_email,omitempty"`
	Mailset             bool   `json:"mailset,omitempty"`
	TrustLevel          int    `json:"trust_level,omitempty"`
	UpdateChannel       bool   `json:"update_channel,omitempty"`
	VolumeCount         int    `json:"volume_count,omitempty"`
	VolumeList          []struct {
		Desc           string `json:"desc,omitempty"`
		Display        string `json:"display,omitempty"`
		MountPoint     string `json:"mount_point,omitempty"`
		SizeFree       string `json:"size_free,omitempty"`
		SizeTotal      string `json:"size_total,omitempty"`
		VolDesc        string `json:"vol_desc,omitempty"`
		VolumeFeatures []any  `json:"volume_features,omitempty"`
	} `json:"volume_list,omitempty"`
}

type PackageInstallUploadRequest struct {
	Additional []string  `form:"additional,json" url:"additional,json"`
	File       form.File `form:"file"                                  kind:"file"`
}

type PackageInstallUploadResponse struct {
	CodesignError int    `json:"codesign_error,omitempty"`
	ID            string `json:"id,omitempty"`
	InstallPages  string `json:"install_pages,omitempty"`
	Licence       string `json:"licence,omitempty"`
	Name          string `json:"name,omitempty"`
	TaskID        string `json:"task_id,omitempty"`
	Version       string `json:"version,omitempty"`
	Additional    struct {
		BreakPkgs            any    `json:"break_pkgs,omitempty"`
		Description          string `json:"description,omitempty"`
		Distributor          string `json:"distributor,omitempty"`
		DsmApps              string `json:"dsm_apps,omitempty"`
		InstallOnColdStorage bool   `json:"install_on_cold_storage,omitempty"`
		InstallReboot        bool   `json:"install_reboot,omitempty"`
		InstallType          string `json:"install_type,omitempty"`
		Maintainer           string `json:"maintainer,omitempty"`
		ReplacePkgs          any    `json:"replace_pkgs,omitempty"`
		Startable            bool   `json:"startable,omitempty"`
		Status               string `json:"status,omitempty"`
		StatusCode           int    `json:"status_code,omitempty"`
		StatusDescription    string `json:"status_description,omitempty"`
		StatusOrigin         string `json:"status_origin,omitempty"`
	} `json:"additional,omitempty"`
}
