package docker

import (
	"strings"
	"time"
)

type Project struct {
	ID                    string    `json:"id,omitempty"`
	Name                  string    `json:"name,omitempty"`
	ContainerIds          []string  `json:"containerIds,omitempty"`
	Containers            []any     `json:"containers,omitempty"`
	Content               string    `json:"content,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitzero"`
	EnableServicePortal   bool      `json:"enable_service_portal,omitempty"`
	IsPackage             bool      `json:"is_package,omitempty"`
	Path                  string    `json:"path,omitempty"`
	ServicePortalName     string    `json:"service_portal_name,omitempty"`
	ServicePortalPort     int       `json:"service_portal_port,omitempty"`
	ServicePortalProtocol string    `json:"service_portal_protocol,omitempty"`
	Services              []struct {
		DisplayName string `json:"display_name,omitempty"`
		ID          string `json:"id,omitempty"`
		ProxyTarget string `json:"proxy_target,omitempty"`
		Service     string `json:"service,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"services,omitempty"`
	SharePath string    `json:"share_path,omitempty"`
	State     string    `json:"state,omitempty"`
	Status    string    `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	Version   int       `json:"version,omitempty"`
}

func (p Project) IsRunning() bool {
	return strings.ToUpper(p.Status) == "RUNNING"
}

type ProjectGetRequest struct {
	ID string `url:"id,omitempty,quoted"`
}

type ProjectListRequest struct {
	Offset int `url:"offset,omitempty"`
	Limit  int `url:"limit,omitempty"`
}

type ProjectListResponse map[string]Project

type ProjectCreateRequest struct {
	Name                  string `url:"name,omitempty,quoted"`
	Content               string `url:"content,omitempty,quoted"`
	SharePath             string `url:"share_path,omitempty,quoted"`
	EnableServicePortal   *bool  `url:"enable_service_portal,omitempty"`
	ServicePortalName     string `url:"service_portal_name,quoted"`
	ServicePortalPort     *int64 `url:"service_portal_port"`
	ServicePortalProtocol string `url:"service_portal_protocol,quoted"`
}

type ProjectCreateResponse struct {
	ContainerIds          []any     `json:"containerIds,omitempty"`
	Containers            []any     `json:"containers,omitempty"`
	Content               string    `json:"content,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
	ID                    string    `json:"id,omitempty"`
	IsPackage             bool      `json:"is_package,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Path                  string    `json:"path,omitempty"`
	EnableServicePortal   *bool     `url:"enable_service_portal,omitempty"`
	ServicePortalName     string    `json:"service_portal_name,omitempty"`
	ServicePortalPort     *int64    `json:"service_portal_port,omitempty"`
	ServicePortalProtocol string    `json:"service_portal_protocol,omitempty"`
	Services              []struct {
		DisplayName string `json:"display_name,omitempty"`
		ID          string `json:"id,omitempty"`
		ProxyTarget string `json:"proxy_target,omitempty"`
		Service     string `json:"service,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"services,omitempty"`
	SharePath string    `json:"share_path,omitempty"`
	State     string    `json:"state,omitempty"`
	Status    string    `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Version   int       `json:"version,omitempty"`
}

type ProjectUpdateRequest struct {
	ID                    string `url:"id,omitempty,quoted"`
	Content               string `url:"content,omitempty,quoted"`
	EnableServicePortal   *bool  `url:"enable_service_portal,omitempty"`
	ServicePortalName     string `url:"service_portal_name,omitempty,quoted"`
	ServicePortalPort     *int64 `url:"service_portal_port,omitempty"`
	ServicePortalProtocol string `url:"service_portal_protocol,omitempty,quoted"`
}

type ProjectUpdateResponse struct {
	ID                    string    `json:"id,omitempty"`
	Name                  string    `json:"name,omitempty"`
	ContainerIds          []string  `json:"containerIds,omitempty"`
	Containers            []any     `json:"containers,omitempty"`
	Content               string    `json:"content,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
	EnableServicePortal   bool      `json:"enable_service_portal,omitempty"`
	IsPackage             bool      `json:"is_package,omitempty"`
	Path                  string    `json:"path,omitempty"`
	ServicePortalName     string    `json:"service_portal_name,omitempty"`
	ServicePortalPort     int       `json:"service_portal_port,omitempty"`
	ServicePortalProtocol string    `json:"service_portal_protocol,omitempty"`
	Services              []struct {
		DisplayName string `json:"display_name,omitempty"`
		ID          string `json:"id,omitempty"`
		ProxyTarget string `json:"proxy_target,omitempty"`
		Service     string `json:"service,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"services,omitempty"`
	SharePath string    `json:"share_path,omitempty"`
	State     string    `json:"state,omitempty"`
	Status    string    `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Version   int       `json:"version,omitempty"`
}

type ProjectDeleteRequest struct {
	ID string `url:"id,omitempty,quoted"`
}

type ProjectDeleteResponse struct{}

type ProjectCleanStreamRequest struct {
	ID string `url:"id,omitempty,quoted"`
}

type ProjectCleanStreamResponse string

type ProjectStreamRequest struct {
	ID string `url:"id,omitempty,quoted"`
}

type ProjectStreamResponse string
