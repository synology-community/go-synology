package webstation

type Portal struct {
	ID                      string `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	ACL                     any    `json:"acl,omitempty"`
	CompatibleCrtService    any    `json:"compatible_crt_service,omitempty"`
	CompatibleCrtSubscriber any    `json:"compatible_crt_subscriber,omitempty"`
	CompatibleScSection     any    `json:"compatible_sc_section,omitempty"`
	Enable                  bool   `json:"enable,omitempty"`
	EnableLog               bool   `json:"enable_log,omitempty"`
	ErrorPage               string `json:"error_page,omitempty"`
	Fqdn                    string `json:"fqdn,omitempty"`
	Hsts                    bool   `json:"hsts,omitempty"`
	HTTPPort                []int  `json:"http_port,omitempty"`
	HTTPSPort               []int  `json:"https_port,omitempty"`
	Preserve                bool   `json:"preserve,omitempty"`
	Service                 string `json:"service,omitempty"`
	Shortcut                struct {
		AllUsers    bool `json:"all_users,omitempty"`
		App         any  `json:"app,omitempty"`
		DisplayName any  `json:"display_name,omitempty"`
		Enable      bool `json:"enable,omitempty"`
		Hide        bool `json:"hide,omitempty"`
		Icon        any  `json:"icon,omitempty"`
	} `json:"shortcut,omitempty"`
	Type string `json:"type,omitempty"`
}

type PortalCreateRequest struct {
	Portal Portal `url:"portal,json" form:"portal,omitempty,json"`
}

type PortalListResponse struct {
	Portals []Portal `json:"portals,omitempty"`
}
