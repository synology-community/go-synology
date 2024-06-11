package api

type ApiInfo = map[string]InfoData
type InfoData struct {
	Path          string `json:"path"`
	MinVersion    int    `json:"min_version"`
	MaxVersion    int    `json:"max_version"`
	RequestFormat string `json:"request_format"`
}
