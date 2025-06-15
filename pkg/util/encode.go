package util

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func EncodeValues(s any, k string, v *url.Values) error {
	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}
	v.Set(k, string(encoded))

	return nil
}

func EncodeValuesWrap(s any, k string, v *url.Values) error {
	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}

	quoted := strconv.Quote(string(encoded))

	v.Set(k, quoted)

	return nil
}
