package docker_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/docker"
)

func newClient(suite *suite.Suite) docker.Api {
	c, err := api.New(api.Options{
		Host:       os.Getenv("SYNOLOGY_HOST"),
		VerifyCert: false,
	})
	suite.Require().NoError(err)

	r, err := c.Login(context.Background(), api.LoginOptions{
		Username: os.Getenv("SYNOLOGY_USER"),
		Password: os.Getenv("SYNOLOGY_PASSWORD"),
	})
	suite.Require().NoError(err)

	suite.T().Log("Login successful")
	suite.T().Logf("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)

	return docker.New(c)
}

type ClientTestSuite struct {
	suite.Suite

	Client docker.Api

	ContainerName   string
	ProjectShare    string
	ProjectName     string
	ImageRepository string
	ImageTag        string
}

func (suite *ClientTestSuite) SetupTest() {
	suite.ContainerName = "test-container"
	suite.ProjectName = "test-project"
	suite.ProjectShare = "/projects"
	suite.ImageRepository = "chainguard/bash"
	suite.ImageTag = "latest"

	client := newClient(&suite.Suite)

	suite.Client = client
}

func (suite *ClientTestSuite) TestImagePull() {
	ctx := context.Background()

	resp, err := suite.Client.ImagePull(ctx, suite.ImageRepository, suite.ImageTag)
	suite.NoError(err)
	suite.Equal(resp.Finished, true)
}

func (suite *ClientTestSuite) TestImageDelete() {
	ctx := context.Background()

	_, err := suite.Client.ImagePull(ctx, suite.ImageRepository, suite.ImageTag)
	if err != nil {
		suite.NoError(err)
		suite.T().Skip("Unable to pull image, skipping delete test")
	}

	resp, err := suite.Client.ImageDelete(ctx, docker.ImageDeleteRequest{
		Images: docker.ImageList{
			{
				Repository: suite.ImageRepository,
				Tags:       []string{suite.ImageTag},
			},
		},
	})
	suite.NoError(err)
	suite.Equal(resp.ImageObjects[suite.ImageRepository][suite.ImageTag].Error, int64(0))
}

func (suite *ClientTestSuite) TestRegistryList() {
	ctx := context.Background()

	resp, err := suite.Client.RegistryList(ctx, docker.ListRegistryRequest{
		Limit:  -1,
		Offset: 0,
	})
	suite.NoError(err)
	suite.Greater(resp.Total, int64(0))
	for _, registry := range resp.Registries {
		suite.NotEmpty(registry.Name)
		suite.NotEmpty(registry.URL)
	}
}

func (suite *ClientTestSuite) TestContainerCreate() {
	ctx := context.Background()

	resp, err := suite.Client.ContainerCreate(ctx, docker.CreateContainerRequest{
		Container: docker.Container{
			Name:  suite.ContainerName,
			Image: suite.ImageRepository + ":" + suite.ImageTag,
			PortBindings: []docker.PortBinding{
				{HostPort: 8080, ContainerPort: 80, Protocol: "tcp"},
			},
		},
	})
	suite.NoError(err)
	suite.NotNil(resp)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func TestNewClient(t *testing.T) {
	type args struct {
		client api.Api
	}
	tests := []struct {
		name string
		args args
		want docker.Api
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := docker.New(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
