package core

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
