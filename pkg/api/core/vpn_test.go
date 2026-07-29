package core

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
