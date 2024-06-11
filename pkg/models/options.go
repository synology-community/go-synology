package models

type RequestOptions struct {
	Host                        string `json:"host"`
	SkipCertificateVerification bool   `json:"skip_certificate_verification"`

	RetryLimit int64 `json:"retry_limit"`
	Timeout    int64 `json:"timeout"`
	Logger     any   `json:"logger"`
}
