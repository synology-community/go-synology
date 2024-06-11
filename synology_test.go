package synology

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/synology-community/go-synology/internal/util/form"
	"github.com/synology-community/go-synology/pkg/api"
)

type Nil struct{}

func newClient(t *testing.T) *Client {
	c, err := New(api.Options{
		Host:       os.Getenv("SYNOLOGY_HOST"),
		VerifyCert: false,
	})
	if err != nil {
		t.Error(err)
		require.NoError(t, err)
	}

	if r, err := c.Login(context.Background(), os.Getenv("SYNOLOGY_USER"), os.Getenv("SYNOLOGY_PASSWORD")); err != nil {
		t.Error(err)
		require.NoError(t, err)
	} else {
		fmt.Printf("Session: %s\nDeviceID: %s\n", r.SessionID, r.DeviceID)
		// log.Infoln("Login successful")
		// log.Infof("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)
	}

	if client, ok := c.(*Client); !ok {
		t.Error("Client is not of type APIClient")
		require.True(t, ok)
		return nil
	} else {
		return client
	}
}

func newSuiteClient(suite *suite.Suite) *Client {
	c, err := New(api.Options{
		Host:       os.Getenv("SYNOLOGY_HOST"),
		VerifyCert: false,
	})
	suite.Require().NoError(err)

	r, err := c.Login(context.Background(), os.Getenv("SYNOLOGY_USER"), os.Getenv("SYNOLOGY_PASSWORD"))
	suite.Require().NoError(err)

	fmt.Printf("Session: %s\nDeviceID: %s\n", r.SessionID, r.DeviceID)
	// log.Infoln("Login successful")
	// log.Infof("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)

	client, ok := c.(*Client)

	suite.Require().True(ok)

	return client
}

func Test_FileStationClient_Upload(t *testing.T) {
	c := newClient(t)

	file := form.File{
		Name:    "test.txt",
		Content: "Hello, World!",
	}

	_, err := c.FileStationAPI().Upload(context.Background(), "/data/foo", file, true, true)
	require.NoError(t, err)
}

func Test_FileStationClient_MD5(t *testing.T) {
	c := newClient(t)

	file := form.File{
		Name:    "test.txt",
		Content: "Hello, World!",
	}

	_, err := c.FileStationAPI().Upload(context.Background(), "/data/foo", file, true, true)
	require.NoError(t, err)

	_, err = c.FileStationAPI().MD5(context.Background(), "/data/foo/test.txt")
	require.NoError(t, err)
}

func Test_FileStationClient_DeleteStart(t *testing.T) {
	c := newClient(t)

	_, err := c.FileStationAPI().DeleteStart(context.Background(), []string{"/data/foodbar"}, true)
	require.NoError(t, err)
}

func Test_FileStationClient_DeleteStatus(t *testing.T) {
	c := newClient(t)

	r, err := c.FileStationAPI().DeleteStart(context.Background(), []string{"/data/foodbar"}, true)
	require.NoError(t, err)

	_, err = c.FileStationAPI().DeleteStatus(context.Background(), r.TaskID)
	require.NoError(t, err)
}
