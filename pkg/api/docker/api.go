package docker

import (
	"context"
)

type Api interface {
	ContainerCreate(
		ctx context.Context,
		container CreateContainerRequest,
	) (*CreateContainerResponse, error)
	ContainerStop(
		ctx context.Context,
		req ContainerOperationRequest,
	) (*ContainerStopResponse, error)
	ContainerStart(
		ctx context.Context,
		req ContainerOperationRequest,
	) (*ContainerStartResponse, error)
	ContainerRestart(
		ctx context.Context,
		req ContainerOperationRequest,
	) (*ContainerRestartResponse, error)

	RegistryList(ctx context.Context, req ListRegistryRequest) (*ListRegistryResponse, error)

	ImagePullStart(ctx context.Context, req ImagePullStartRequest) (*ImagePullStartResponse, error)
	ImagePullStatus(
		ctx context.Context,
		req ImagePullStatusRequest,
	) (*ImagePullStatusResponse, error)
	ImagePull(ctx context.Context, repository, tag string) (*ImagePullStatusResponse, error)
	ImageDelete(ctx context.Context, req ImageDeleteRequest) (*ImageDeleteResponse, error)

	ProjectGet(ctx context.Context, id string) (*Project, error)
	ProjectGetByName(ctx context.Context, name string) (*Project, error)
	ProjectList(ctx context.Context, req ProjectListRequest) ([]Project, error)
	ProjectCreate(ctx context.Context, req ProjectCreateRequest) (*ProjectCreateResponse, error)
	ProjectUpdate(ctx context.Context, req ProjectUpdateRequest) (*ProjectUpdateResponse, error)
	ProjectDelete(ctx context.Context, req ProjectDeleteRequest) (*ProjectDeleteResponse, error)
	ProjectCleanStream(
		ctx context.Context,
		req ProjectStreamRequest,
	) (*ProjectStreamResponse, error)
	ProjectStopStream(ctx context.Context, req ProjectStreamRequest) (*ProjectStreamResponse, error)
	ProjectRestartStream(
		ctx context.Context,
		req ProjectStreamRequest,
	) (*ProjectStreamResponse, error)
	ProjectStartStream(
		ctx context.Context,
		req ProjectStreamRequest,
	) (*ProjectStreamResponse, error)
	ProjectBuildStream(
		ctx context.Context,
		req ProjectStreamRequest,
	) (*ProjectStreamResponse, error)

	NetworkList(ctx context.Context) ([]Network, error)
	NetworkGetByName(ctx context.Context, name string) (*Network, error)
	NetworkGetByID(ctx context.Context, id string) (*Network, error)
	NetworkCreate(ctx context.Context, req Network) error
	NetworkUpdate(ctx context.Context, req NetworkUpdateRequest) error
	NetworkDelete(ctx context.Context, networks ...Network) error
}
