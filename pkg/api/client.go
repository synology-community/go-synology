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
	"github.com/pquerna/otp/totp"
	"github.com/synology-community/go-synology/pkg/query"
	"github.com/synology-community/go-synology/pkg/util"
	"github.com/synology-community/go-synology/pkg/util/form"
	"golang.org/x/net/publicsuffix"
)

const (
	API_BASE = "/webapi/entry.cgi"
)

var defaultTimeout = 30 * time.Second

type Client struct {
	httpClient *retryablehttp.Client

	BaseURL *url.URL

	ApiCredentials Credentials

	username string
	password string

	// once sync.Once
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
	baseURL.Path = API_BASE

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

// Session data to bypass login process.
type Session struct {
	SessionID string    `json:"sid"`
	SynoToken string    `json:"syno_token"`
	CreatedAt time.Time `json:"created_at"`
}

// ExportSession returns the current session structure (SID, token and timestamp).
func (c *Client) ExportSession() Session {
	return Session{
		SessionID: c.ApiCredentials.SessionID,
		SynoToken: c.ApiCredentials.Token,
		CreatedAt: time.Now(),
	}
}

// ImportSession injects a previously exported session into the client and wires it into the base URL.
func (c *Client) ImportSession(s Session) {
	c.ApiCredentials = Credentials{
		SessionID: s.SessionID,
		Token:     s.SynoToken,
	}
	if c.BaseURL != nil {
		q := c.BaseURL.Query()
		if s.SessionID != "" {
			q.Set("_sid", s.SessionID)
		}
		if s.SynoToken != "" {
			q.Set("SynoToken", s.SynoToken)
		}
		c.BaseURL.RawQuery = q.Encode()
	}
}

// IsSessionAlive verifies if current session (usually imported) is still valid by calling a info API.
func (c *Client) IsSessionAlive(ctx context.Context) (bool, error) {
	if c.BaseURL == nil {
		return false, errors.New("base url is nil")
	}
	if c.ApiCredentials.SessionID == "" && c.ApiCredentials.Token == "" {
		return false, nil
	}
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	if data, err := c.GetUserInfo(ctx); err == nil {
		if data.UserName == "" {
			return false, fmt.Errorf("session probe failed to query user info")
		}
		return true, nil
	} else {
		var ae ApiError
		if errors.As(err, &ae) {
			if ae.Code == 119 {
				return false, nil
			}
			return false, fmt.Errorf("session probe failed with code %d", ae.Code)
		}
		return false, err
	}
}

var (
	ErrOtpRejected = errors.New(
		"OTP code rejected (invalid or reused)",
	) // special case to handle retries
	ErrOtpRequired = errors.New(
		"OTP code is required by the server, but was not provided (password was correct)",
	)
)

// Login runs a login flow to retrieve session token from Synology.
func (c *Client) Login(ctx context.Context, options LoginOptions) (*LoginResponse, error) {
	username := options.Username
	password := options.Password
	otpSecret := options.OTPSecret
	var token, sessionID, deviceID string

	c.username = username
	c.password = password

	now := time.Now()

	req := LoginRequest{
		Account:         username,
		Password:        password,
		Format:          "sid",
		EnableSynoToken: "yes",
		TimeZone:        now.Format("-07:00"),
	}

	if otpSecret != "" {
		otpCode, err := totp.GenerateCode(otpSecret, now.Local())
		if err == nil {
			req.OTPCode = otpCode
		} else {
			return nil, multierror.Append(err, errors.New("unable to generate otp code"))
		}
	}

	resp, err := Get[LoginResponse](c, ctx, &req, Login)
	if err != nil {
		if terr, ok := err.(PermissionDeniedError); ok {
			tmpToken, err := terr.GetToken()
			if err != nil {
				return nil, multierror.Append(err, errors.New("unable to get token from error"))
			}
			req.OTPCode, err = totp.GenerateCode(otpSecret, now.Local())
			if err != nil {
				return nil, multierror.Append(err, errors.New("unable to generate otp code"))
			}
			req.Password = tmpToken
			resp, err = Get[LoginResponse](c, ctx, &req, Login)
			if err != nil {
				return nil, multierror.Append(
					err,
					fmt.Errorf("unable to login with token: %v", tmpToken),
				)
			} else {
				if resp.Token != "" {
					token = resp.Token
				}
				if resp.SessionID != "" {
					sessionID = resp.SessionID
				}
				if resp.DeviceID != "" {
					deviceID = resp.DeviceID
				}
			}
		} else {
			if apiErr, ok := err.(ApiError); ok && apiErr.Code == 404 {
				return nil, ErrOtpRejected
			}
			if otpSecret != "" {
				return nil, multierror.Append(err, errors.New("unable to login using TOTP and password"))
			} else {
				return nil, multierror.Append(err, errors.New("unable to login using password"))
			}
		}
	} else {
		if resp.Token != "" {
			token = resp.Token
		}
		if resp.SessionID != "" {
			sessionID = resp.SessionID
		}
		if resp.DeviceID != "" {
			deviceID = resp.DeviceID
		}
	}

	if token != "" && sessionID != "" {
		c.ApiCredentials = Credentials{
			SessionID: sessionID,
			Token:     token,
			DeviceID:  deviceID,
		}
		q := c.BaseURL.Query()
		q.Set("_sid", sessionID)
		q.Set("SynoToken", token)

		c.BaseURL.RawQuery = q.Encode()

		return resp, nil
	} else {
		return resp, errors.New("unable to login")
	}
}

func PostFileUpload[TResp Response](
	c Api,
	ctx context.Context,
	name string,
	content string,
	method Method,
) (*TResp, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	defer func() { _ = w.Close() }()

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

	return Do[TResp](c.Client(), req, method.ErrorSummaries)
}

func mergeQueries(qs ...any) (url.Values, error) {
	res := map[string][]string{}
	for _, q := range qs {
		qq, err := query.Values(q)
		if err != nil {
			return nil, err
		}
		maps.Copy(res, qq)
	}
	return res, nil
}

func getQuery(c Api, p ...any) (string, error) {
	if c.BaseUrl() == nil {
		return "", errors.New("base url is nil")
	}

	ps := make([]any, 0, len(p)+1)
	ps = append(ps, c.BaseUrl().Query())
	ps = append(ps, p...)

	q, err := mergeQueries(ps)
	if err != nil {
		return "", err
	}

	return q.Encode(), nil
}

func PostFileWithQuery[TResp Response, TReq Request](
	c Api,
	ctx context.Context,
	r *TReq,
	method Method,
) (*TResp, error) {
	q, err := getQuery(c, method, r)
	if err != nil {
		return nil, err
	}
	u := *c.BaseUrl()
	u.RawQuery = q

	return postFile[TResp](c.Client(), ctx, u.String(), r)
}

func PostFile[TResp Response, TReq Request](
	c Api,
	ctx context.Context,
	r *TReq,
	method Method,
) (*TResp, error) {
	return postFile[TResp](c.Client(), ctx, c.BaseUrl().String(), method, r)
}

func postFile[TResp Response](
	c *retryablehttp.Client,
	ctx context.Context,
	url string,
	input ...any,
) (*TResp, error) {
	var errorSummaries ErrorSummaries

	if method, ok := input[0].(Method); !ok {
		errorSummaries = GlobalErrors
	} else {
		errorSummaries = method.ErrorSummaries
	}

	buf := new(bytes.Buffer)
	if w, fs, err := form.Marshal(buf, input...); err != nil {
		_ = w.Close()
		return nil, err
	} else {
		defer func() { _ = w.Close() }()

		// Only set a timeout if one isn't already set
		var cancel context.CancelFunc
		if _, ok := ctx.Deadline(); !ok {
			ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
		}

		req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, url, buf)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Length", fmt.Sprintf("%d", fs))
		req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", w.Boundary()))

		return Do[TResp](c, req, errorSummaries)
	}
}

func List[T Response](c Api, ctx context.Context, method Method) (*T, error) {
	return Get[T, Request](c, ctx, nil, method)
}

func Void[TReq Request](c Api, ctx context.Context, r *TReq, method Method) error {
	_, err := Get[Request](c, ctx, r, method)
	return err
}

func GetEncode[TResp Response, TReq EncodeRequest](
	c Api,
	ctx context.Context,
	r *TReq,
	method Method,
) (*TResp, error) {
	return Get[TResp](c, ctx, r, method)
}

func Post[TResp Response, TReq Request](
	c Api,
	ctx context.Context,
	r *TReq,
	method Method,
) (*TResp, error) {
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

	req, err := retryablehttp.NewRequestWithContext(
		ctx,
		http.MethodPost,
		u.String(),
		strings.NewReader(qu.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if c.Credentials().Token != "" {
		req.Header.Set("X-SYNO-TOKEN", c.Credentials().Token)
		req.Header.Set("Cookie", c.Credentials().GetCookie())
	}

	return Do[TResp](c.Client(), req, method.ErrorSummaries)
}

func GetQuery[TResp any](c Api, ctx context.Context, r any, method Method) (*TResp, error) {
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
	maps.Copy(aq, qu)
	maps.Copy(aq, dq)

	u := c.BaseUrl()

	u.RawQuery = aq.Encode()

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

	return Do[TResp](c.Client(), req, method.ErrorSummaries)
}

func Get[TResp Response, TReq Request](
	c Api,
	ctx context.Context,
	r *TReq,
	method Method,
) (*TResp, error) {
	// return GetQuery[TResp](c, ctx, r, method)
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

	return Do[TResp](c.Client(), req, method.ErrorSummaries)
}

func download(r io.ReadCloser) (any, error) {
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

func Do[T Response](
	client *retryablehttp.Client,
	req *retryablehttp.Request,
	errorSummaries ErrorSummaries,
) (*T, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}()

	return handle[T](resp, errorSummaries)
}

func handle[T Response](resp *http.Response, errorSummaries ErrorSummaries) (*T, error) {
	var synoResponse ApiResponse[T]
	var synoResponsePartialAuth ApiResponsePartialAuth[T]

	contentType := resp.Header.Get("Content-Type")
	contentType = strings.Split(contentType, ";")[0]

	switch contentType {
	case "application/json":
		if respBody, readErr := io.ReadAll(resp.Body); readErr == nil {
			if decodeErr := json.NewDecoder(bytes.NewReader(respBody)).Decode(&synoResponse); decodeErr != nil {
				if decodeErr := json.NewDecoder(bytes.NewReader(respBody)).Decode(&synoResponsePartialAuth); decodeErr == nil {
					return nil, ErrOtpRequired
				} else {
					return nil, errors.New("unable to decode response: " + decodeErr.Error() + "\n\n" + string(respBody))
				}
			}
		} else {
			return nil, readErr
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
		return nil, handleErrors(synoResponse, errorSummaries)
	}
}

func handleErrors[T Response](response ApiResponse[T], knownErrors ErrorSummaries) error {
	if response.Error.Code == 0 {
		return nil
	}

	return response.Error.WithSummaries(knownErrors)
}

type NotFoundError ApiError

func (e NotFoundError) Error() string {
	msg := "not found"
	if e.Summary != "" {
		msg = e.Summary
	}
	_ = multierror.Append(fmt.Errorf("%s", msg), e)
	return msg
}

type PermissionDeniedError ApiError

func (e PermissionDeniedError) Error() string {
	msg := "permission denied"
	if e.Summary != "" {
		msg = e.Summary
	}
	_ = multierror.Append(fmt.Errorf("%s", msg), e)
	return msg
}

func (e PermissionDeniedError) GetToken() (token string, err error) {
	if len(e.Errors) > 0 {
		for _, fields := range e.Errors {
			if f, ok := fields.Fields["token"]; ok {
				if t, ok := f.(string); ok {
					token = t
				} else {
					err = errors.New("unable to parse token")
				}
			} else {
				err = errors.New("token not found")
			}
		}
	}
	return
}
