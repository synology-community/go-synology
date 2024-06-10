package models

import (
	"fmt"
	"net/url"
	"strings"
)

type JsonArray []string

func (s JsonArray) EncodeValues(k string, v *url.Values) error {
	v.Set(k, fmt.Sprintf(`["%s"]`, strings.Join(s, `","`)))
	return nil
}

type JsonString string

func (s JsonString) EncodeValues(k string, v *url.Values) error {
	v.Set(k, `"`+string(s)+`"`)
	return nil
}
