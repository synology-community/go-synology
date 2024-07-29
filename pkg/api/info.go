package api

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed api_info.json
var b []byte

var ApiInfoData ApiInfo = unmarshalApiInfoData()

func unmarshalApiInfoData() ApiInfo {
	var data ApiInfo
	err := json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type ApiInfo = map[string]InfoData
type InfoData struct {
	Path          string `json:"path"`
	MinVersion    int    `json:"minVersion"`
	MaxVersion    int    `json:"maxVersion"`
	RequestFormat string `json:"requestFormat"`
}

func (c *Client) GetApiInfo(ctx context.Context) (*map[string]InfoData, error) {
	return List[ApiInfo](c, ctx, Api_Info)
}
