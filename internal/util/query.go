package util

import (
	"maps"
	"net/url"

	"github.com/google/go-querystring/query"
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
