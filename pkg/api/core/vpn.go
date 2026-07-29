package core

import (
	"context"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/core/methods"
	"github.com/synology-community/go-synology/pkg/util/form"
)

// OpenVPNClientProfileRequest creates an OpenVPN client connection profile on
// the DSM VPN Client from an imported .ovpn configuration file.
//
// Endpoint: SYNO.Core.Network.VPN.OpenVPNWithConf, method create. This API is
// undocumented; its field names and multipart/form-data shape were captured
// from the request DSM's own "Create VPN Profile" wizard sends (it is not a
// regular JSON-parameter API call like most SYNO.Core.* endpoints).
type OpenVPNClientProfileRequest struct {
	ConfigName     string    `form:"confname"`
	Server         string    `form:"server"`
	User           string    `form:"user"`
	Pass           string    `form:"pass"`
	PresharedKey   string    `form:"preshared_key"`
	Port           string    `form:"port"`
	Protocol       string    `form:"protocol"`
	OVPNFile       form.File `form:"ovpn_file"       kind:"file"`
	CAFile         form.File `form:"ca_file"         kind:"file"`
	ClientCrtFile  form.File `form:"client_crt_file" kind:"file"`
	ClientKeyFile  form.File `form:"client_key_file" kind:"file"`
	PemFile        form.File `form:"pem_file"        kind:"file"`
	TAFile         form.File `form:"ta_file"         kind:"file"`
	Compress       string    `form:"compress"`
	DefaultGateway bool      `form:"defgw"`
	NAT            bool      `form:"nat"`
	Reconnect      bool      `form:"reconnect"`
}

type OpenVPNClientProfileResponse struct{}

func (c *Client) OpenVPNClientProfileCreate(
	ctx context.Context,
	req OpenVPNClientProfileRequest,
) (*OpenVPNClientProfileResponse, error) {
	return api.PostFileWithQuery[OpenVPNClientProfileResponse](
		c.client,
		ctx,
		&req,
		methods.NetworkVPNOpenVPNWithConfCreate,
	)
}
