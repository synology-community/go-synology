package docker

type Network struct {
	ID                string   `json:"id,omitempty"                 url:"-"`
	Name              string   `json:"name,omitempty"               url:"name"`
	Subnet            string   `json:"subnet,omitempty"             url:"subnet,omitempty"`
	IPRange           string   `json:"iprange,omitempty"            url:"iprange,omitempty"`
	Gateway           string   `json:"gateway,omitempty"            url:"gateway,omitempty"`
	DisableMasquerade bool     `json:"disable_masquerade,omitempty" url:"disable_masquerade"`
	EnableIPv6        bool     `json:"enable_ipv6,omitempty"        url:"enable_ipv6"`
	EnableIPv4        bool     `json:"enable_ipv4,omitempty"        url:"enable_ipv4"`
	Driver            string   `json:"driver,omitempty"             url:"driver"`
	Auto              bool     `json:"auto,omitempty"               url:"auto"`
	Containers        []string `json:"containers,omitempty"`
}

// NetworkCreateRequest represents a network creation request.
type NetworkCreateRequest struct {
	Name              string `json:"name"                         url:"name"`
	Driver            string `json:"driver,omitempty"             url:"driver"`
	Subnet            string `json:"subnet,omitempty"             url:"subnet,omitempty"`
	IPRange           string `json:"iprange,omitempty"            url:"iprange,omitempty"`
	Gateway           string `json:"gateway,omitempty"            url:"gateway,omitempty"`
	EnableIPv4        bool   `json:"enable_ipv4,omitempty"        url:"enable_ipv4"`
	EnableIPv6        bool   `json:"enable_ipv6,omitempty"        url:"enable_ipv6"`
	DisableMasquerade bool   `json:"disable_masquerade,omitempty" url:"disable_masquerade"`
	Auto              bool   `json:"auto,omitempty"               url:"auto"`
}

// NetworkCreateResponse represents a network creation response.
type NetworkCreateResponse struct {
	Success bool `json:"success"`
}

// NetworkListRequest represents a network list request.
type NetworkListRequest struct {
	Additional []string `json:"additional,omitempty" url:"additional"`
}

// NetworkListResponse represents a network list response.
type NetworkListResponse struct {
	Data struct {
		Network []Network `json:"network"`
	} `json:"data"`
	Success bool `json:"success"`
}

// NetworkDeleteRequest represents a network deletion request.
type NetworkDeleteRequest struct {
	ID   string `json:"id,omitempty"   url:"id"`
	Name string `json:"name,omitempty" url:"name"`
}

// NetworkDeleteResponse represents a network deletion response.
type NetworkDeleteResponse struct {
	Success bool `json:"success"`
}

// NetworkGetRequest represents a network get request.
type NetworkGetRequest struct {
	ID   string `json:"id,omitempty"   url:"id"`
	Name string `json:"name,omitempty" url:"name"`
}

// NetworkGetResponse represents a network get response.
type NetworkGetResponse struct {
	Data    Network `json:"data"`
	Success bool    `json:"success"`
}
