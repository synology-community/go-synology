package models

type RequestOptions struct {
	RetryLimit int64 `json:"retry_limit"`
	Timeout    int64 `json:"timeout"`
}
