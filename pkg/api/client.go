package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/synology-community/go-synology/pkg/query"
	"github.com/synology-community/go-synology/pkg/util"
	"github.com/synology-community/go-synology/pkg/util/form"
	"golang.org/x/net/publicsuffix"
)

var defaultTimeout = 15 * time.Second

type Client struct {
	httpClient *retryablehttp.Client

	BaseURL *url.URL

	ApiCredentials Credentials

	username string
	password string
}

func New(o Options) (Api, error) {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !o.VerifyCert,
		},
	}

	// currently, 'Cookie' is the only supported method for providing 'sid' token to DSM
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	c := retryablehttp.NewClient()
	if o.Logger != nil {
		c.Logger = o.Logger
	}
	c.HTTPClient.Jar = jar
	c.HTTPClient.Transport = transport

	baseURL, err := url.Parse(o.Host)

	baseURL.Scheme = "https"
	baseURL.Path = "/webapi/entry.cgi"

	if err != nil {
		return nil, err
	}

	client := &Client{
		httpClient: c,
		BaseURL:    baseURL,
	}

	return client, nil
}

func (c *Client) Password() string {
	return c.password
}

// BaseUrl implements api.Client.
func (c *Client) BaseUrl() *url.URL {
	return c.BaseURL
}

// Client implements api.Client.
func (c *Client) Client() *retryablehttp.Client {
	return c.httpClient
}

func (c *Client) Credentials() Credentials {
	return c.ApiCredentials
}

// Login runs a login flow to retrieve session token from Synology.
func (c *Client) Login(ctx context.Context, user, password, otpSecret string) (*LoginResponse, error) {

	c.username = user
	c.password = password

	req := LoginRequest{
		Account:  user,
		Password: password,
		// Session:         sessionName,
		Format:          "sid", //"cookie",
		EnableSynoToken: "yes",
	}

	if otpSecret != "" {
		otpCode, err := generateTotp(otpSecret)
		if err != nil {
			return nil, err
		}
		req.OTPCode = otpCode
	}

	resp, err := Get[LoginResponse](c, ctx, &req, Login)
	if err != nil {
		return nil, err
	}
	c.ApiCredentials = Credentials{
		SessionID: resp.SessionID,
		Token:     resp.Token,
	}
	q := c.BaseURL.Query()
	q.Set("_sid", resp.SessionID)
	q.Set("SynoToken", resp.Token)

	c.BaseURL.RawQuery = q.Encode()
	return resp, nil
}

func PostFileUpload[TResp Response](c Api, ctx context.Context, name string, content string, method Method) (*TResp, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	defer w.Close()

	fs := int64(0)

	fileReader := strings.NewReader(content)

	if fw, err := w.CreateFormFile("file", name); err != nil {
		return nil, err
	} else {

		if size, err := io.Copy(fw, fileReader); err != nil {
			return nil, err
		} else {
			fs = size
		}

	}

	u := c.BaseUrl()
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Length", fmt.Sprintf("%d", fs))
	req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", w.Boundary()))

	return Do[TResp](c.Client(), req)
}

func PostFile[TResp Response, TReq Request](c Api, ctx context.Context, r *TReq, method Method) (*TResp, error) {
	buf := new(bytes.Buffer)

	// Prepare a form that you will submit to that URL.
	if w, fs, err := form.Marshal(buf, method, r); err != nil {
		w.Close()
		return nil, err
	} else {
		defer w.Close()

		// Only set a timeout if one isn't already set
		var cancel context.CancelFunc
		if _, ok := ctx.Deadline(); !ok {
			ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
		}

		u := c.BaseUrl()
		req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, u.String(), buf)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Length", fmt.Sprintf("%d", fs))
		req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", w.Boundary()))

		return Do[TResp](c.Client(), req)
	}
}

func List[T Response](c Api, ctx context.Context, method Method) (*T, error) {
	return Get[T, Request](c, ctx, nil, method)
}

func Void[TReq Request](c Api, ctx context.Context, r *TReq, method Method) error {
	_, err := Get[Request](c, ctx, r, method)
	return err
}

func GetEncode[TResp Response, TReq EncodeRequest](c Api, ctx context.Context, r *TReq, method Method) (*TResp, error) {
	return Get[TResp](c, ctx, r, method)
}

func Post[TResp Response, TReq Request](c Api, ctx context.Context, r *TReq, method Method) (*TResp, error) {
	qu, err := util.Query(method, r, c.Credentials())
	if err != nil {
		return nil, err
	}

	u := c.BaseUrl().JoinPath(method.API)
	u.RawQuery = ""

	// Only set a timeout if one isn't already set
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, u.String(), strings.NewReader(qu.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-SYNO-TOKEN", c.Credentials().Token)

	return Do[TResp](c.Client(), req)
}

func GetQuery[TResp any](c Api, ctx context.Context, r interface{}, method Method) (*TResp, error) {
	aq, err := query.Values(method) //.AsApiParams())
	if err != nil {
		return nil, err
	}
	dq, err := query.Values(r)
	if err != nil {
		return nil, err
	}

	url := c.BaseUrl()

	qu := maps.Clone(url.Query())
	maps.Copy(qu, aq)
	maps.Copy(qu, dq)

	u := c.BaseUrl()

	u.RawQuery = qu.Encode()

	// Only set a timeout if one isn't already set
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return Do[TResp](c.Client(), req)
}

func Get[TResp Response, TReq Request](c Api, ctx context.Context, r *TReq, method Method) (*TResp, error) {
	aq, err := query.Values(method) //.AsApiParams())
	if err != nil {
		return nil, err
	}
	dq, err := query.Values(r)
	if err != nil {
		return nil, err
	}

	u2 := c.BaseUrl()

	qu := maps.Clone(u2.Query())
	maps.Copy(qu, aq)
	maps.Copy(qu, dq)

	if u2 == nil {
		return nil, errors.New("base url is nil")
	}
	u := new(url.URL)
	*u = *u2

	u.RawQuery = qu.Encode()

	// Only set a timeout if one isn't already set
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return Do[TResp](c.Client(), req)
}

func download(r io.ReadCloser) (interface{}, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return &form.File{
		Content: buf.String(),
		Name:    "download",
	}, nil
}

func Do[T Response](client *retryablehttp.Client, req *retryablehttp.Request) (*T, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}()

	return handleResponse[T](resp)
}

func handle[T Response](resp *http.Response, res *T) error {
	r, err := handleResponse[T](resp)
	if err != nil {
		return err
	}
	*res = *r
	return nil
}

func handleResponse[T Response](resp *http.Response) (*T, error) {
	var synoResponse ApiResponse[T]

	contentType := resp.Header.Get("Content-Type")
	contentType = strings.Split(contentType, ";")[0]

	switch contentType {
	case "application/json":
		if err := json.NewDecoder(resp.Body).Decode(&synoResponse); err != nil {
			return nil, err
		}
	case "application/octet-stream":
		resp, err := download(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp, ok := resp.(*T); ok {
			return resp, nil
		} else {
			return nil, errors.New("invalid response")
		}
	}

	if synoResponse.Success {
		return &synoResponse.Data, nil
	} else {
		return nil, handleErrors(synoResponse, GlobalErrors)
	}
}

func handleErrors[T Response](response ApiResponse[T], knownErrors ErrorSummaries) error {
	if response.Error.Code == 0 {
		return nil
	}

	var result error

	if errDesc, ok := knownErrors()[response.Error.Code]; ok {
		result = multierror.Append(result, fmt.Errorf("api response error code %d: %v", response.Error.Code, errDesc))
	} else {
		result = multierror.Append(result, fmt.Errorf("api response error code %d: %v", response.Error.Code, response.Error))
	}

	if response.Error.Errors != nil {
		for i, err := range response.Error.Errors {
			if errDesc, ok := knownErrors()[err.Code]; ok {
				result = multierror.Append(result, fmt.Errorf("api response error[%d] code %d: %v", i, err.Code, errDesc))
			} else {
				result = multierror.Append(result, fmt.Errorf("api response error[%d] code %d: %v", i, err.Code, err))
			}
		}
	}

	return result
}
