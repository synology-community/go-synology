package webstation

type Service struct {
	ID                   string   `json:"id,omitempty"`
	Category             string   `json:"category,omitempty"`
	ConnectTimeout       int      `json:"connect_timeout,omitempty"`
	CustomRule           struct{} `json:"custom_rule,omitempty"`
	Description          any      `json:"description,omitempty"`
	DisplayName          string   `json:"display_name,omitempty"`
	DisplayNameI18N      any      `json:"display_name_i18n,omitempty"`
	Enable               bool     `json:"enable,omitempty"`
	Icon                 string   `json:"icon,omitempty"`
	ManagedByDocker      bool     `json:"managed_by_docker,omitempty"`
	ManagedByWebService  bool     `json:"managed_by_web_service,omitempty"`
	ProxyHeaders         []any    `json:"proxy_headers,omitempty"`
	ProxyHTTPVersion     int      `json:"proxy_http_version,omitempty"`
	ProxyInterceptErrors bool     `json:"proxy_intercept_errors,omitempty"`
	ProxyTarget          string   `json:"proxy_target,omitempty"`
	ReadTimeout          int      `json:"read_timeout,omitempty"`
	Root                 string   `json:"root,omitempty"`
	SendTimeout          int      `json:"send_timeout,omitempty"`
	Service              string   `json:"service,omitempty"`
	Status               struct {
		ErrCode int    `json:"err_code,omitempty"`
		Type    string `json:"type,omitempty"`
	} `json:"status,omitempty"`
	SupportAlias  bool   `json:"support_alias,omitempty"`
	SupportServer bool   `json:"support_server,omitempty"`
	Type          string `json:"type,omitempty"`
}

type ServiceListResponse struct {
	Services []Service `json:"services,omitempty"`
}
