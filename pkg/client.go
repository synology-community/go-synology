package client

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/api/virtualization"
	"golang.org/x/net/publicsuffix"
)

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
	httpClient *retryablehttp.Client

	FileStation    *fileStationClient
	Virtualization *virtualizationClient

	BaseURL url.URL

	Auth AuthStorage
}

// Login runs a login flow to retrieve session token from Synology.
func (c *APIClient) Login(ctx context.Context, user, password string) (*api.LoginResponse, error) {
	resp, err := Get[api.LoginRequest, api.LoginResponse](c, ctx, &api.LoginRequest{
		Account:  user,
		Password: password,
		// Session:         sessionName,
		Format:          "sid", //"cookie",
		EnableSynoToken: "yes",
	}, api.API_METHODS["Login"])
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
func New(o Options) (SynologyClient, error) {
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

	synoClient := &APIClient{
		httpClient: c,
		BaseURL:    *baseURL,
	}

	synoClient.FileStation = &fileStationClient{client: synoClient}
	synoClient.Virtualization = &virtualizationClient{client: synoClient}

	return synoClient, nil
}
