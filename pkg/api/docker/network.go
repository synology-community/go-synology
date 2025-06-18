package docker

type Network struct {
	ID                string   `json:"id,omitempty"                 url:"-"`
	Name              string   `json:"name"                         url:"name,quoted"`
	Subnet            string   `json:"subnet,omitempty"             url:"subnet,quoted,omitempty"`
	IPRange           string   `json:"iprange,omitempty"            url:"iprange,quoted,omitempty"`
	Gateway           string   `json:"gateway,omitempty"            url:"gateway,quoted,omitempty"`
	DisableMasquerade bool     `json:"disable_masquerade,omitempty" url:"disable_masquerade,omitempty"`
	EnableIPv6        bool     `json:"enable_ipv6"                  url:"enable_ipv6"`
	IPv6Subnet        string   `json:"ipv6_subnet,omitempty"        url:"ipv6_subnet,quoted,omitempty"`
	IPv6Gateway       string   `json:"ipv6_gateway,omitempty"       url:"ipv6_gateway,quoted,omitempty"`
	IPv6IPRange       string   `json:"ipv6_iprange,omitempty"       url:"ipv6_iprange,quoted,omitempty"`
	Driver            string   `json:"driver"                       url:"driver"`
	Containers        []string `json:"containers,omitempty"         url:"containers,json,omitempty"`
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
	Network []Network `json:"network,omitempty"`
}

// NetworkDeleteRequest represents a network deletion request.
type NetworkDeleteRequest struct {
	Networks []Network `json:"networks,omitempty" url:"networks,json"`
}

// NetworkDeleteResponse represents a network deletion response.
type NetworkDeleteResponse struct {
	Failed []string `json:"failed,omitempty"`
}

type NetworkUpdateRequest struct {
	Name       string   `url:"networkName,quoted"`
	Containers []string `url:"containers,json"`
}

type NetworkUpdateResponse struct {
	AddFailed       []string `json:"add_failed,omitempty"`
	AddSucceeded    []string `json:"add_succeeded,omitempty"`
	RemoveFailed    []string `json:"remove_failed,omitempty"`
	RemoveSucceeded []string `json:"remove_succeeded,omitempty"`
}
