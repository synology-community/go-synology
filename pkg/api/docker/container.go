package docker

import (
	"net/url"

	"github.com/synology-community/go-synology/pkg/util"
)

type EnvVariable struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Link struct {
	LinkContainer string `json:"link_container,omitempty"`
	Alias         string `json:"alias,omitempty"`
}

type ContainerNetwork struct {
	Name   string `json:"name,omitempty"`
	Driver string `json:"driver,omitempty"`
}

type PortBinding struct {
	HostPort      int64  `json:"host_port,omitempty"`
	ContainerPort int64  `json:"container_port,omitempty"`
	Protocol      string `json:"type,omitempty"`
}

type VolumeBinding struct {
	HostVolumeFile string `json:"host_volume_file,omitempty"`
	MountPoint     string `json:"mount_point,omitempty"`
	Type           string `json:"type,omitempty"`
}

type Container struct {
	Name                string             `json:"name,omitempty"`
	Image               string             `json:"image,omitempty"`
	Privileged          bool               `json:"privileged,omitempty"`
	PortBindings        []PortBinding      `json:"port_bindings,omitempty"`
	VolumeBindings      []VolumeBinding    `json:"volume_bindings,omitempty"`
	EnvVariables        []EnvVariable      `json:"env_variables,omitempty"`
	Network             []ContainerNetwork `json:"network,omitempty"`
	UseHostNetwork      bool               `json:"use_host_network,omitempty"`
	Cmd                 string             `json:"cmd,omitempty"`
	ServicePortals      []string           `json:"service_portals,omitempty"`
	CPUPriority         int64              `json:"cpu_priority,omitempty"`
	MemoryLimit         int64              `json:"memory_limit,omitempty"`
	EnableRestartPolicy bool               `json:"enable_restart_policy,omitempty"`
	CapAdd              []string           `json:"CapAdd,omitempty"`
	CapDrop             []string           `json:"CapDrop,omitempty"`
	Links               []Link             `json:"links,omitempty"`
}

func (s Container) EncodeValues(k string, v *url.Values) error {
	return util.EncodeValues(&s, k, v)
}

type CreateContainerRequest struct {
	Container      Container `json:"profile,omitempty"          url:"profile"`
	IsRunInstantly bool      `json:"is_run_instantly,omitempty" url:"is_run_instantly"`
}

type CreateContainerResponse struct {
	Services                []string `json:"services,omitempty"`
	StartDependentContainer bool     `json:"start_dependent_container,omitempty"`
}

type ContainerOperationRequest struct {
	Name string `json:"name,omitempty" url:"name,omitempty,quoted"`
}

type ContainerOperationResponse struct {
	CPU           float64 `json:"cpu,omitempty"`
	Memory        int64   `json:"memory,omitempty"`
	MemoryPercent float64 `json:"memoryPercent,omitempty"`
	Name          string  `json:"name,omitempty"`
}

type (
	ContainerStopResponse = ContainerOperationResponse
)

type (
	ContainerStartResponse struct {
		ContainerOperationResponse
		StartDependentContainer bool `json:"start_dependent_container,omitempty"`
	}
)

type (
	ContainerRestartResponse = ContainerOperationResponse
)
