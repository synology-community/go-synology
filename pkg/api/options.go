package api

type Options struct {
	Host            string `json:"host"`
	VerifyCert      bool   `json:"skip_certificate_verification"`
	AllowHTTP       bool   `json:"allow_http"`

	RetryLimit int64 `json:"retry_limit"`
	Timeout    int64 `json:"timeout"`
	Logger     any   `json:"logger"`
}
