package virtualization

import (
	"net/url"
	"strings"
)

type Storage struct {
	ID         string `url:"storage_id"            json:"storage_id"`
	Name       string `url:"storage_name"          json:"storage_name"`
	Status     string `url:"status"                json:"status"`
	HostName   string `url:"host_name,omitempty"   json:"host_name,omitempty"`
	HostID     string `url:"host_id,omitempty"     json:"host_id,omitempty"`
	Size       int    `url:"size,omitempty"        json:"size,omitempty"`
	Used       int    `url:"used,omitempty"        json:"used,omitempty"`
	VolumePath string `url:"volume_path,omitempty" json:"volume_path,omitempty"`
}

type StorageList struct {
	Storages []Storage `json:"storages"`
}

type Storages []Storage

func (s Storages) EncodeValues(k string, v *url.Values) error {
	var storageIDs, storageNames []string

	for _, storage := range s {
		if storage.ID != "" {
			storageIDs = append(storageIDs, storage.ID)
		}
		if storage.Name != "" {
			storageNames = append(storageNames, storage.Name)
		}
	}

	if len(storageIDs) > 0 {
		encStorageIDs := `["` + strings.Join(storageIDs, `","`) + `"]`
		v.Set("storage_ids", encStorageIDs)
	}

	if len(storageNames) > 0 {
		encStorageNames := `["` + strings.Join(storageNames, `","`) + `"]`
		v.Set("storage_names", encStorageNames)
	}

	return nil
}
