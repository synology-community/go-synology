package virtualization

import "net/url"

type Guest struct {
	ID          string `mapstructure:"guest_id" json:"guest_id"`
	Name        string `mapstructure:"guest_name" json:"guest_name"`
	Description string `mapstructure:"description" json:"description"`
	Status      string `mapstructure:"status" json:"status"`
	StorageID   string `mapstructure:"storage_id" json:"storage_id"`
	StorageName string `mapstructure:"storage_name" json:"storage_name"`
	Autorun     int    `mapstructure:"autorun" json:"autorun"`
	VcpuNum     int    `mapstructure:"vcpu_num" json:"vcpu_num"`
	VramSize    int    `mapstructure:"vram_size" json:"vram_size"`
	Disks       []struct {
		Controller int    `mapstructure:"controller" json:"controller"`
		Unmap      bool   `mapstructure:"unmap" json:"unmap"`
		ID         string `mapstructure:"vdisk_id" json:"vdisk_id"`
		Size       int    `mapstructure:"vdisk_size" json:"vdisk_size"`
	} `mapstructure:"vdisks" json:"vdisks"`
	Networks []struct {
		ID     string `mapstructure:"network_id" json:"network_id"`
		Name   string `mapstructure:"network_name" json:"network_name"`
		Mac    string `mapstructure:"mac" json:"mac"`
		Model  int    `mapstructure:"model" json:"model"`
		VnicID string `mapstructure:"vnic_id" json:"vnic_id"`
	} `mapstructure:"vnics" json:"vnics"`
}

type GuestList struct {
	Guests []Guest `mapstructure:"guests" json:"guests"`
}

type GetGuest struct {
	Name string `form:"guest_name" url:"guest_name"`
}

func (l GetGuest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}
