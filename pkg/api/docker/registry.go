package docker

type Registry struct {
	EnableRegistryMirror bool     `json:"enable_registry_mirror,omitempty"`
	EnableTrustSSC       bool     `json:"enable_trust_SSC,omitempty"`
	MirrorURLs           []string `json:"mirror_urls,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Syno                 bool     `json:"syno,omitempty"`
	URL                  string   `json:"url,omitempty"`
}

type ListRegistryResponse struct {
	Registries []Registry `json:"registries,omitempty"`
	Offset     int64      `json:"offset,omitempty"`
	Total      int64      `json:"total,omitempty"`
	Using      string     `json:"using,omitempty"`
}

type ListRegistryRequest struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}
