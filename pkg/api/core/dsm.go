package core

import "github.com/synology-community/go-synology/pkg/models"

type PortEnableRequest struct {
	IsDirectID bool             `url:"isDirectID"`
	IsPkg      bool             `url:"isPkg"`
	Name       models.JsonArray `url:"name"`
}

type PortEnableResponse struct {
	IsPortAllow bool `json:"isPortAllow"`
	PortCheck   bool `json:"portCheck"`
}
