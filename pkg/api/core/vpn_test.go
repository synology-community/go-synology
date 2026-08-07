package core

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/go-synology/pkg/query"
	"github.com/synology-community/go-synology/pkg/util/form"
)

func TestOpenVPNClientProfileRequest_Marshal(t *testing.T) {
	req := OpenVPNClientProfileRequest{
		ConfigName:     "pornic",
		User:           "freebox-pornic",
		Pass:           "s3cr3t",
		PresharedKey:   "s3cr3t",
		Port:           "1194",
		Protocol:       "udp",
		OVPNFile:       form.File{Name: "perreux.ovpn", Content: "client\nremote 1.2.3.4 1194\n"},
		DefaultGateway: true,
		NAT:            true,
		Reconnect:      true,
	}

	buf := new(bytes.Buffer)
	w, _, err := form.Marshal(buf, &req)
	require.NoError(t, err)
	require.NoError(t, w.Close())

	r := multipart.NewReader(buf, w.Boundary())

	fields := map[string]string{}
	files := map[string]string{}
	for {
		part, err := r.NextPart()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)

		name := part.FormName()
		content, err := io.ReadAll(part)
		require.NoError(t, err)

		if strings.Contains(part.Header.Get("Content-Disposition"), "filename=") {
			files[name] = string(content)
		} else {
			fields[name] = string(content)
		}
	}

	require.Equal(t, "pornic", fields["confname"])
	require.Equal(t, "", fields["server"])
	require.Equal(t, "freebox-pornic", fields["user"])
	require.Equal(t, "s3cr3t", fields["pass"])
	require.Equal(t, "s3cr3t", fields["preshared_key"])
	require.Equal(t, "1194", fields["port"])
	require.Equal(t, "udp", fields["protocol"])
	require.Equal(t, "true", fields["defgw"])
	require.Equal(t, "true", fields["nat"])
	require.Equal(t, "true", fields["reconnect"])

	require.Equal(t, "client\nremote 1.2.3.4 1194\n", files["ovpn_file"])
	// Unused file slots are still sent as empty parts, matching what DSM's own
	// "Create VPN Profile" wizard sends.
	require.Contains(t, files, "ca_file")
	require.Contains(t, files, "client_crt_file")
	require.Contains(t, files, "client_key_file")
	require.Contains(t, files, "pem_file")
	require.Contains(t, files, "ta_file")
}

func TestOpenVPNClientProfileListRequest_EncodeValues(t *testing.T) {
	req := OpenVPNClientProfileListRequest{Additional: []string{"status"}}

	v, err := query.Values(&req)
	require.NoError(t, err)

	require.Contains(t, v.Encode(), "additional=%5B%22status%22%5D")
}

func TestOpenVPNClientProfileDeleteRequest_EncodeValues(t *testing.T) {
	req := OpenVPNClientProfileDeleteRequest{ID: "o1786051381"}

	v, err := query.Values(&req)
	require.NoError(t, err)

	require.Equal(t, "id=%22o1786051381%22", v.Encode())
}

func TestOpenVPNClientProfileListResponse_Unmarshal(t *testing.T) {
	// Captured from the "data" field of a real
	// SYNO.Core.Network.VPN.OpenVPNWithConf list response; the outer
	// {"data": ..., "success": true} envelope is unwrapped by the client's
	// handle() before decoding into TResp, so only the inner array is
	// exercised here.
	body := []byte(
		`[{"confname":"perreux","defgw":false,"id":"o1786140031","nat":true,"pass":"\t\t\t\t\t\t\t\t","prtl":"ovpn_conf","reconnect":true,"rx":"0","server":"82.64.94.63","status":"connected","tx":"0","uptime":"114","user":"perreux","virtual_ip":"192.168.27.66","vpn_gateway":"212.27.38.253"}]`,
	)

	var resp OpenVPNClientProfileListResponse
	require.NoError(t, json.Unmarshal(body, &resp))

	require.Len(t, resp, 1)

	profile := resp[0]
	require.Equal(t, "o1786140031", profile.ID)
	require.Equal(t, "perreux", profile.ConfigName)
	require.Equal(t, "perreux", profile.User)
	require.Equal(t, "82.64.94.63", profile.Server)
	require.False(t, profile.DefaultGateway)
	require.True(t, profile.NAT)
	require.True(t, profile.Reconnect)
	require.Equal(t, "ovpn_conf", profile.Protocol)
	require.Equal(t, "connected", profile.Status)
}
