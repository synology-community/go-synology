package filestation

import (
	"net/url"
)

type MD5StartRequest struct {
	Path string `form:"file_path" url:"file_path"`
}

type MD5StartResponse struct {
	TaskID string `json:"taskid"`
}

type UrlWrapString string

func (s UrlWrapString) EncodeValues(k string, v *url.Values) error {

	v.Set(k, `"`+string(s)+`"`)

	return nil
}

type MD5StatusRequest struct {
	TaskID UrlWrapString `url:"taskid" form:"taskid"`
}

type MD5StatusResponse struct {
	Finished bool   `json:"finished"`
	MD5      string `json:"md5"`
}

type MD5Response struct {
	MD5 string `json:"md5"`
}
