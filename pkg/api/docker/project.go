package docker

import "time"

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
	ID string `json:"id,omitempty"`
}

type ProjectGetResponse Project

type ProjectListRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type ProjectListResponse map[string]Project

type ProjectCreateRequest struct {
	Name                  string `json:"name,omitempty"`
	Content               string `json:"content,omitempty"`
	SharePath             string `json:"share_path,omitempty"`
	EnableServicePortal   bool   `json:"enable_service_portal,omitempty"`
	ServicePortalName     string `json:"service_portal_name,omitempty"`
	ServicePortalPort     int    `json:"service_portal_port,omitempty"`
	ServicePortalProtocol string `json:"service_portal_protocol,omitempty"`
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
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
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
	ID string `json:"id,omitempty"`
}

type ProjectDeleteResponse struct{}

type ProjectCleanStreamRequest struct {
	ID string `json:"id,omitempty"`
}

type ProjectCleanStreamResponse string
