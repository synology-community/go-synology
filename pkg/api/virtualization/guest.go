package virtualization

import "net/url"

type Guest struct {
	ID          string `url:"guest_id" json:"guest_id"`
	Name        string `url:"guest_name" json:"guest_name"`
	Description string `url:"description" json:"description"`
	Status      string `url:"status" json:"status"`
	StorageID   string `url:"storage_id" json:"storage_id"`
	StorageName string `url:"storage_name" json:"storage_name"`
	Autorun     int    `url:"autorun" json:"autorun"`
	VcpuNum     int    `url:"vcpu_num" json:"vcpu_num"`
	VramSize    int    `url:"vram_size" json:"vram_size"`
	Disks       []struct {
		Controller int    `url:"controller" json:"controller"`
		Unmap      bool   `url:"unmap" json:"unmap"`
		ID         string `url:"vdisk_id" json:"vdisk_id"`
		Size       int    `url:"vdisk_size" json:"vdisk_size"`
	} `url:"vdisks" json:"vdisks"`
	Networks []struct {
		ID     string `url:"network_id" json:"network_id"`
		Name   string `url:"network_name" json:"network_name"`
		Mac    string `url:"mac" json:"mac"`
		Model  int    `url:"model" json:"model"`
		VnicID string `url:"vnic_id" json:"vnic_id"`
	} `url:"vnics" json:"vnics"`
}

type GuestList struct {
	Guests []Guest `url:"guests" json:"guests"`
}

type GetGuest struct {
	Name string `form:"guest_name" url:"guest_name"`
}

func (l GetGuest) EncodeValues(_ string, _ *url.Values) error {
	return nil
}
