package filestation

import (
	"github.com/synology-community/synology-api/package/api"
)

type FileStationInfoRequest api.ApiParams

type FileStationInfoResponse struct {
	IsManager              bool
	SupportVirtualProtocol string
	Supportsharing         bool
	Hostname               string
}
