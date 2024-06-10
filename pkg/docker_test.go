package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synology-community/synology-api/pkg/api/docker"
)

type DockerClientTestSuite struct {
	suite.Suite

	Client docker.DockerApi

	ContainerName   string
	ProjectShare    string
	ProjectName     string
	ImageRepository string
	ImageTag        string
}

func (suite *DockerClientTestSuite) SetupTest() {
	suite.ContainerName = "test-container"
	suite.ProjectName = "test-project"
	suite.ProjectShare = "/projects"
	suite.ImageRepository = "chainguard/bash"
	suite.ImageTag = "latest"

	client := newSuiteClient(&suite.Suite)

	suite.Client = client.Docker
}

func (suite *DockerClientTestSuite) TestImagePull() {
	ctx := context.Background()

	resp, err := suite.Client.ImagePull(ctx, suite.ImageRepository, suite.ImageTag)
	suite.NoError(err)
	suite.Equal(resp.Finished, true)
}

func (suite *DockerClientTestSuite) TestImageDelete() {
	ctx := context.Background()

	_, err := suite.Client.ImagePull(ctx, suite.ImageRepository, suite.ImageTag)
	if err != nil {
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

func (suite *DockerClientTestSuite) TestRegistryList() {
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

func (suite *DockerClientTestSuite) TestContainerCreate() {
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

func TestDockerClientTestSuite(t *testing.T) {
	suite.Run(t, new(DockerClientTestSuite))
}

func TestNewDockerClient(t *testing.T) {
	type args struct {
		client *APIClient
	}
	tests := []struct {
		name string
		args args
		want docker.DockerApi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDockerClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDockerClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
