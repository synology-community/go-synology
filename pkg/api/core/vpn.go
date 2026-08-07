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

// DSM's create response carries no useful data (not even the new profile's
// id) — confirmed empirically, so the id must be discovered afterwards via
// OpenVPNClientProfileList.
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

// OpenVPNClientProfile is a single profile entry as returned by
// SYNO.Core.Network.VPN.OpenVPNWithConf, method list.
type OpenVPNClientProfile struct {
	ID             string `json:"id"`
	ConfigName     string `json:"confname"`
	User           string `json:"user"`
	Server         string `json:"server"`
	DefaultGateway bool   `json:"defgw"`
	NAT            bool   `json:"nat"`
	Reconnect      bool   `json:"reconnect"`
	Protocol       string `json:"prtl"`
	Status         string `json:"status,omitempty"`
}

type OpenVPNClientProfileListRequest struct {
	Additional []string `url:"additional,omitempty,json"`
}

type OpenVPNClientProfileListResponse []OpenVPNClientProfile

// OpenVPNClientProfileList lists all OpenVPN client profiles configured on
// the DSM VPN Client. Use this to discover a profile's id (needed for
// OpenVPNClientProfileDelete), since create doesn't return it.
func (c *Client) OpenVPNClientProfileList(
	ctx context.Context,
) (*OpenVPNClientProfileListResponse, error) {
	return api.Get[OpenVPNClientProfileListResponse](
		c.client,
		ctx,
		&OpenVPNClientProfileListRequest{Additional: []string{"status"}},
		methods.NetworkVPNOpenVPNWithConfList,
	)
}

// OpenVPNClientProfileDeleteRequest deletes a profile by id (e.g.
// "o1786051381"), as returned by OpenVPNClientProfileList.
type OpenVPNClientProfileDeleteRequest struct {
	ID string `url:"id,quoted"`
}

// OpenVPNClientProfileDelete deletes an existing OpenVPN client profile.
// Endpoint: SYNO.Core.Network.VPN.OpenVPNWithConf, method delete. Like
// create, this is undocumented; captured from DSM's own VPN Client UI.
func (c *Client) OpenVPNClientProfileDelete(ctx context.Context, id string) error {
	return api.Void(
		c.client,
		ctx,
		&OpenVPNClientProfileDeleteRequest{ID: id},
		methods.NetworkVPNOpenVPNWithConfDelete,
	)
}
