package docker

import (
	"time"

	"github.com/synology-community/synology-api/pkg/models"
)

type Project struct {
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

type ProjectGetRequest struct {
	ID models.JsonString `url:"id,omitempty"`
}

type ProjectGetResponse Project

type ProjectListRequest struct {
	Offset int `url:"offset,omitempty"`
	Limit  int `url:"limit,omitempty"`
}

type ProjectListResponse map[string]Project

type ProjectCreateRequest struct {
	Name                  models.JsonString `url:"name,omitempty"`
	Content               models.JsonString `url:"content,omitempty"`
	SharePath             models.JsonString `url:"share_path,omitempty"`
	EnableServicePortal   bool              `url:"enable_service_portal,omitempty"`
	ServicePortalName     models.JsonString `url:"service_portal_name"`
	ServicePortalPort     int               `url:"service_portal_port"`
	ServicePortalProtocol models.JsonString `url:"service_portal_protocol"`
}

type ProjectCreateResponse struct {
	ContainerIds          []any     `json:"containerIds,omitempty"`
	Containers            []any     `json:"containers,omitempty"`
	Content               string    `json:"content,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
	EnableServicePortal   bool      `json:"enable_service_portal,omitempty"`
	ID                    string    `json:"id,omitempty"`
	IsPackage             bool      `json:"is_package,omitempty"`
	Name                  string    `json:"name,omitempty"`
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

type ProjectUpdateRequest struct {
	ID      models.JsonString `url:"id,omitempty"`
	Content models.JsonString `url:"content,omitempty"`
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
	ID models.JsonString `url:"id,omitempty"`
}

type ProjectDeleteResponse struct{}

type ProjectCleanStreamRequest struct {
	ID models.JsonString `url:"id,omitempty"`
}

type ProjectCleanStreamResponse string
