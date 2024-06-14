package util

import (
	"maps"
	"net/url"

	"github.com/synology-community/go-synology/pkg/query"
)

func Query(params ...any) (url.Values, error) {
	result := url.Values{}

	for _, param := range params {
		c, err := query.Values(param)
		if err != nil {
			return nil, err
		}
		maps.Copy(result, c)
	}

	return result, nil
}
