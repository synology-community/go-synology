package core

// Group represents a Synology DSM group.
type Group struct {
	ID          string `json:"gid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// GroupListRequest for listing groups.
type GroupListRequest struct {
	Additional []string `url:"additional,omitempty"`
}

// GroupListResponse for listing groups.
type GroupListResponse struct {
	Groups []Group `json:"groups,omitempty"`
	Offset int     `json:"offset,omitempty"`
	Total  int     `json:"total,omitempty"`
}

// GroupCreateRequest for creating a group.
type GroupCreateRequest struct {
	Name        string `url:"name"`
	Description string `url:"description,omitempty"`
}

// GroupCreateResponse for creating a group.
type GroupCreateResponse struct {
	Group Group `json:"group,omitempty"`
}

// GroupDeleteRequest for deleting a group.
type GroupDeleteRequest struct {
	Name string `url:"name"`
}

// GroupDeleteResponse for deleting a group.
type GroupDeleteResponse struct {
	Success bool `json:"success"`
}

// GroupModifyRequest for modifying a group.
type GroupModifyRequest struct {
	Name        string `url:"name"`
	NewName     string `url:"new_name,omitempty"`
	Description string `url:"description,omitempty"`
}

// GroupModifyResponse for modifying a group.
type GroupModifyResponse struct {
	Group Group `json:"group,omitempty"`
}
