package util

import (
	"encoding/json"
	"net/url"
	"strings"
)

func EncodeValues(s interface{}, k string, v *url.Values) error {

	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}
	v.Set(k, string(encoded))

	return nil
}

func EncodeValuesWrap(s interface{}, k string, v *url.Values) error {

	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}

	e := strings.ReplaceAll(string(encoded), `"`, `\"`)

	v.Set(k, `"`+e+`"`)

	return nil
}
