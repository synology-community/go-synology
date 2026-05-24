package core

type BackendProxy struct {
	Fqdn     string `json:"fqdn"`
	Port     int64  `json:"port"`
	Protocol int64  `json:"protocol"`
}

type FrontendProxy struct {
	Acl   *string `json:"acl"`
	Fqdn  string  `json:"fqdn"`
	Port  int64   `json:"port"`
	Https struct {
		Hsts bool `json:"hsts"`
	} `json:"https"`
	Protocol int64 `json:"protocol"`
}

type CustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ReverseProxy struct {
	UUID                 string         `json:"UUID"`
	Key                  string         `json:"_key"`
	Backend              BackendProxy   `json:"backend,omitempty"`
	CustomizeHeaders     []CustomHeader `json:"customize_headers,omitempty"`
	Description          string         `json:"description,omitempty"`
	Frontend             FrontendProxy  `json:"frontend,omitempty"`
	ProxyConnectTimeout  int64          `json:"proxy_connect_timeout"`
	ProxyHttpVersion     int64          `json:"proxy_http_version,omitempty"`
	ProxyInterceptErrors bool           `json:"proxy_intercept_errors,omitempty"`
	ProxyReadTimeout     int64          `json:"proxy_read_timeout,omitempty"`
	ProxySendTimeout     int64          `json:"proxy_send_timeout,omitempty"`
}

type ReverseProxyCreate struct {
	Backend              map[string]any `json:"backend,omitempty"`
	CustomizeHeaders     []CustomHeader `json:"customize_headers"`
	Description          string         `json:"description,omitempty"`
	Frontend             map[string]any `json:"frontend,omitempty"`
	ProxyConnectTimeout  int64          `json:"proxy_connect_timeout"`
	ProxyHttpVersion     int64          `json:"proxy_http_version,omitempty"`
	ProxyInterceptErrors bool           `json:"proxy_intercept_errors"`
	ProxyReadTimeout     int64          `json:"proxy_read_timeout,omitempty"`
	ProxySendTimeout     int64          `json:"proxy_send_timeout,omitempty"`
}
type ReverseProxyModify struct {
	ReverseProxyCreate
	UUID string `json:"UUID"`
	Key  string `json:"_key"`
}

type ReverseProxyListRequest struct {
}

type ReverseProxyListResponse struct {
	Entries []ReverseProxy `json:"entries,omitempty"`
}

type ReverseProxyCreateRequest struct {
	Entry      ReverseProxyCreate `url:"entry,json"`
	Additional []string           `url:"additional,omitempty"`
}

// ReverseProxyCreateResponse for creating a reverse proxy entry
type ReverseProxyCreateResponse struct {
}

// ReverseProxyDeleteRequest for deleting a reverse proxy entry
type ReverseProxyDeleteRequest struct {
	UUIDs []string `url:"uuids,json"`
}

// ReverseProxyDeleteResponse for deleting a reverse proxy entry
type ReverseProxyDeleteResponse struct {
}

type ReverseProxyModifyRequest struct {
	Entry ReverseProxyModify `url:"entry,json"`
}

// ReverseProxyModifyResponse for modifying a reverse proxy entry.
type ReverseProxyModifyResponse struct {
}
