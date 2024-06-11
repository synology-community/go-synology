package util

import (
	"encoding/json"
	"net/url"
)

func EncodeValues(s interface{}, k string, v *url.Values) error {

	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}
	v.Set(k, string(encoded))

	return nil
}
