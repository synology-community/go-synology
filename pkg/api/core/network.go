package core

// NetworkConfig represents the network configuration structure.
type NetworkConfig struct {
	ArpIgnore              bool        `json:"arp_ignore"`
	DNSManual              bool        `json:"dns_manual"`
	DNSPrimary             string      `json:"dns_primary"`
	DNSSecondary           string      `json:"dns_secondary"`
	EnableIPConflictDetect bool        `json:"enable_ip_conflict_detect"`
	EnableWindomain        bool        `json:"enable_windomain"`
	Gateway                string      `json:"gateway"`
	GatewayInfo            GatewayInfo `json:"gateway_info"`
	IPv4First              bool        `json:"ipv4_first"`
	MultiGateway           bool        `json:"multi_gateway"`
	ServerName             string      `json:"server_name"`
	UseDHCPDomain          bool        `json:"use_dhcp_domain"`
	V6Gateway              string      `json:"v6gateway"`
}

// GatewayInfo represents the gateway information structure.
type GatewayInfo struct {
	Interface string `json:"ifname"`
	IP        string `json:"ip"`
	Mask      string `json:"mask"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UseDHCP   bool   `json:"use_dhcp"`
}
