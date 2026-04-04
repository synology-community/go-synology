package virtualization

type Network struct {
	ID   string `json:"network_id"`
	Name string `json:"network_name"`
}

type NetworkList struct {
	Networks []Network `json:"networks"`
}
