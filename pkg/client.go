package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/api/virtualization"
	"golang.org/x/net/publicsuffix"
)

var MethodLookup = map[string]api.APIMethodLookup{
	"API":            api.API_METHODS,
	"Virtualization": virtualization.API_METHODS,
	"FileStation":    filestation.API_METHODS,
}

func GetMethod(api, method string) (*api.APIMethod, error) {
	if res, ok := MethodLookup[api][method]; ok {
		return &res, nil
	} else {
		return nil, fmt.Errorf("method not found: %s.%s", api, method)
	}
}

type AuthStorage struct {
	SessionID string `url:"_sid"`
	Token     string `url:"SynoToken"`
}

type SynologyClient interface {
	api.API

	VirtualizationAPI() virtualization.VirtualizationAPI

	FileStationAPI() filestation.FileStationApi

	// get(request api.Request, response api.Response) error
}
type APIClient struct {
	httpClient *http.Client

	FileStation    *fileStationClient
	Virtualization *virtualizationClient

	BaseURL url.URL

	Auth AuthStorage
}

// Login runs a login flow to retrieve session token from Synology.
func (c *APIClient) Login(user, password string) (*api.LoginResponse, error) {
	resp, err := Get[api.LoginRequest, api.LoginResponse](c, &api.LoginRequest{
		Account:  user,
		Password: password,
		// Session:         sessionName,
		Format:          "sid", //"cookie",
		EnableSynoToken: "yes",
	}, api.APIMethod{
		API:          "SYNO.API.Auth",
		Version:      7,
		Method:       "login",
		ErrorSummary: api.GlobalErrors,
	})
	if err != nil {
		return nil, err
	}
	c.Auth = AuthStorage{
		SessionID: resp.SessionID,
		Token:     resp.Token,
	}
	q := c.BaseURL.Query()
	q.Set("_sid", resp.SessionID)
	q.Set("SynoToken", resp.Token)

	c.BaseURL.RawQuery = q.Encode()
	return resp, nil
}

// FileStationAPI implements SynologyClient.
func (c *APIClient) FileStationAPI() filestation.FileStationApi {
	return c.FileStation
}

func (c *APIClient) VirtualizationAPI() virtualization.VirtualizationAPI {
	return c.Virtualization
}

// New initializes "client" instance with minimal input configuration.
func New(host string, skipCertificateVerification bool) (SynologyClient, error) {
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
			InsecureSkipVerify: skipCertificateVerification,
		},
	}

	// currently, 'Cookie' is the only supported method for providing 'sid' token to DSM
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Transport: transport,
		Jar:       jar,
	}

	baseURL, err := url.Parse(host)

	baseURL.Scheme = "https"
	baseURL.Path = "/webapi/entry.cgi"

	if err != nil {
		return nil, err
	}

	synoClient := &APIClient{
		httpClient: httpClient,
		BaseURL:    *baseURL,
	}

	synoClient.FileStation = &fileStationClient{client: synoClient}
	synoClient.Virtualization = &virtualizationClient{client: synoClient}

	return synoClient, nil
}
