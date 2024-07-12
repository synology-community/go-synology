package docker

type Network struct {
	ID                string `json:"id,omitempty" url:"-"`
	Name              string `json:"name,omitempty" url:"name"`
	Subnet            string `json:"subnet,omitempty" url:"subnet,omitempty"`
	IPRange           string `json:"iprange,omitempty" url:"iprange,omitempty"`
	Gateway           string `json:"gateway,omitempty" url:"gateway,omitempty"`
	DisableMasquerade bool   `json:"disable_masquerade,omitempty" url:"disable_masquerade"`
	EnableIPv6        bool   `json:"enable_ipv6,omitempty" url:"enable_ipv6"`
}
