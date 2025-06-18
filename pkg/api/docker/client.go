package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/docker/methods"
	"golang.org/x/exp/maps"
)

type Client struct {
	client api.Api
}

// ProjectCleanStream implements DockerApi.
func (d *Client) ProjectCleanStream(
	ctx context.Context,
	req ProjectStreamRequest,
) (*ProjectStreamResponse, error) {
	return api.Post[ProjectStreamResponse](d.client, ctx, &req, methods.ProjectCleanStream)
}

// ProjectStopStream implements DockerApi.
func (d *Client) ProjectStopStream(
	ctx context.Context,
	req ProjectStreamRequest,
) (*ProjectStreamResponse, error) {
	return api.Post[ProjectStreamResponse](d.client, ctx, &req, methods.ProjectStopStream)
}

// ProjectStartStream implements DockerApi.
func (d *Client) ProjectStartStream(
	ctx context.Context,
	req ProjectStreamRequest,
) (*ProjectStreamResponse, error) {
	return api.Post[ProjectStreamResponse](d.client, ctx, &req, methods.ProjectStartStream)
}

// ProjectBuildStream implements DockerApi.
func (d *Client) ProjectBuildStream(
	ctx context.Context,
	req ProjectStreamRequest,
) (*ProjectStreamResponse, error) {
	return api.Post[ProjectStreamResponse](d.client, ctx, &req, methods.ProjectBuildStream)
}

// ProjectCreate implements DockerApi.
func (d *Client) ProjectCreate(
	ctx context.Context,
	req ProjectCreateRequest,
) (*ProjectCreateResponse, error) {
	return api.Post[ProjectCreateResponse](d.client, ctx, &req, methods.ProjectCreate)
}

// ProjectDelete implements DockerApi.
func (d *Client) ProjectDelete(
	ctx context.Context,
	req ProjectDeleteRequest,
) (*ProjectDeleteResponse, error) {
	return api.Post[ProjectDeleteResponse](d.client, ctx, &req, methods.ProjectDelete)
}

// ProjectGet implements DockerApi.
func (d *Client) ProjectGet(ctx context.Context, id string) (*Project, error) {
	return api.Post[Project](d.client, ctx, &ProjectGetRequest{
		ID: id,
	}, methods.ProjectGet)
}

type ProjectNotFoundError struct{}

func (e ProjectNotFoundError) Error() string {
	return "Project not found"
}

func (d *Client) ProjectGetByName(ctx context.Context, name string) (*Project, error) {
	res, err := d.ProjectList(ctx, ProjectListRequest{
		Offset: 0,
		Limit:  -1,
	})
	if err != nil {
		return nil, err
	}

	for _, p := range res {
		if p.Name == name {
			return &p, nil
		}
	}
	return nil, ProjectNotFoundError{}
}

// ProjectList implements DockerApi.
func (d *Client) ProjectList(ctx context.Context, req ProjectListRequest) ([]Project, error) {
	resp, err := api.Post[map[string]Project](d.client, ctx, &req, methods.ProjectList)
	if err != nil {
		return nil, err
	}
	return maps.Values(*resp), nil
}

// ProjectUpdate implements DockerApi.
func (d *Client) ProjectUpdate(
	ctx context.Context,
	req ProjectUpdateRequest,
) (*ProjectUpdateResponse, error) {
	return api.Post[ProjectUpdateResponse](d.client, ctx, &req, methods.ProjectUpdate)
}

// ImageDelete implements DockerApi.
func (d *Client) ImageDelete(
	ctx context.Context,
	req ImageDeleteRequest,
) (*ImageDeleteResponse, error) {
	return api.Post[ImageDeleteResponse](d.client, ctx, &req, methods.ImageDelete)
}

// ImagePull implements DockerApi.
func (d *Client) ImagePull(
	ctx context.Context,
	repository string,
	tag string,
) (*ImagePullStatusResponse, error) {
	res, err := d.ImagePullStart(ctx, ImagePullStartRequest{
		Repository: repository,
		Tag:        tag,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to delete file, got error: %s", err)
	}

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(60 * time.Second)
	}

	delay := 5 * time.Second
	for {
		task, err := d.ImagePullStatus(ctx, ImagePullStatusRequest{
			TaskID: res.TaskID,
		})
		if err != nil && task == nil {
			return nil, err
		}
		if task.Finished {
			return task, nil
		}
		if time.Now().After(deadline.Add(delay)) {
			return nil, fmt.Errorf("timeout waiting for task to complete")
		}
		time.Sleep(delay)
	}
}

// ImagePullStart implements DockerApi.
func (d *Client) ImagePullStart(
	ctx context.Context,
	req ImagePullStartRequest,
) (*ImagePullStartResponse, error) {
	return api.Post[ImagePullStartResponse](d.client, ctx, &req, methods.ImagePullStart)
}

// ImagePullStatus implements DockerApi.
func (d *Client) ImagePullStatus(
	ctx context.Context,
	req ImagePullStatusRequest,
) (*ImagePullStatusResponse, error) {
	return api.Post[ImagePullStatusResponse](d.client, ctx, &req, methods.ImagePullStatus)
}

// ContainerCreate implements DockerApi.
func (d *Client) ContainerCreate(
	ctx context.Context,
	req CreateContainerRequest,
) (*CreateContainerResponse, error) {
	return api.Post[CreateContainerResponse](d.client, ctx, &req, methods.Create)
}

// RegistryList implements DockerApi.
func (d *Client) RegistryList(
	ctx context.Context,
	req ListRegistryRequest,
) (*ListRegistryResponse, error) {
	return api.Post[ListRegistryResponse](d.client, ctx, &req, methods.RegistryList)
}

// NetworkList implements DockerApi.
func (d *Client) NetworkList(
	ctx context.Context,
) ([]Network, error) {
	resp, err := api.List[NetworkListResponse](d.client, ctx, methods.NetworkList)
	if err != nil {
		return nil, fmt.Errorf("unable to list networks, got error: %s", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("no networks found")
	}
	return resp.Network, nil
}

// NetworkGetByName is a convenience method to get a network by name.
func (d *Client) NetworkGetByName(ctx context.Context, name string) (*Network, error) {
	resp, err := d.NetworkList(ctx)
	if err != nil {
		return nil, err
	}

	for _, network := range resp {
		if network.Name == name {
			return &network, nil
		}
	}
	return nil, fmt.Errorf("network with name '%s' not found", name)
}

// NetworkGetByID is a convenience method to get a network by ID.
func (d *Client) NetworkGetByID(ctx context.Context, id string) (*Network, error) {
	resp, err := d.NetworkList(ctx)
	if err != nil {
		return nil, err
	}

	for _, network := range resp {
		if network.ID == id {
			return &network, nil
		}
	}
	return nil, fmt.Errorf("network with ID '%s' not found", id)
}

// NetworkCreate implements DockerApi.
func (d *Client) NetworkCreate(
	ctx context.Context,
	req Network,
) error {
	_, err := api.Post[api.Data](d.client, ctx, &req, methods.NetworkCreate)
	if err != nil {
		return fmt.Errorf("unable to create network, got error: %s", err)
	}
	return nil
}

// NetworkUpdate implements DockerApi.
func (d *Client) NetworkUpdate(
	ctx context.Context,
	req NetworkUpdateRequest,
) error {
	resp, err := api.Post[NetworkUpdateResponse](d.client, ctx, &req, methods.NetworkSet)
	if err != nil {
		return fmt.Errorf("unable to update network, got error: %s", err)
	}
	if len(resp.AddFailed) > 0 || len(resp.RemoveFailed) > 0 {
		return fmt.Errorf("Network update failed: add_failed=%v, remove_failed=%v",
			resp.AddFailed, resp.RemoveFailed)
	}
	return nil
}

// NetworkDelete implements DockerApi.
func (d *Client) NetworkDelete(
	ctx context.Context,
	networks ...Network,
) error {
	resp, err := api.Post[NetworkDeleteResponse](d.client, ctx, &NetworkDeleteRequest{
		Networks: networks,
	}, methods.NetworkDelete)
	if err != nil {
		return fmt.Errorf("unable to delete network, got error: %s", err)
	}
	if len(resp.Failed) > 0 {
		return fmt.Errorf("Network deletion failed: %v", resp.Failed)
	}
	return nil
}

func New(client api.Api) Api {
	return &Client{client: client}
}
