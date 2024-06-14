package core

type PortEnableRequest struct {
	IsDirectID bool     `url:"isDirectID"`
	IsPkg      bool     `url:"isPkg"`
	Name       []string `url:"name,json"`
}

type PortEnableResponse struct {
	IsPortAllow bool `json:"isPortAllow"`
	PortCheck   bool `json:"portCheck"`
}
