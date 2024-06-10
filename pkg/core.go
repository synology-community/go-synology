package client

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"maps"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/google/go-querystring/query"
// 	"github.com/hashicorp/go-multierror"
// 	"github.com/hashicorp/go-retryablehttp"
// 	"github.com/synology-community/synology-api/pkg/api"
// 	"github.com/synology-community/synology-api/pkg/api/core"
// 	"github.com/synology-community/synology-api/pkg/api/core/methods"
// 	"github.com/synology-community/synology-api/pkg/util"
// 	"github.com/synology-community/synology-api/pkg/util/form"
// )

// var defaultTimeout = 15 * time.Second

// type coreClient struct {
// 	client *APIClient
// }

// // SystemInfo implements core.CoreApi.
// func (c *coreClient) SystemInfo(ctx context.Context) (*core.SystemInfoResponse, error) {
// 	panic("unimplemented")
// }

// // PackageList implements core.CoreApi.
// func (c *coreClient) PackageList(ctx context.Context) (*core.PackageListResponse, error) {
// 	return List[core.PackageListResponse](c.client, ctx, methods.PackageList)
// }

// func NewCoreClient(client *APIClient) core.CoreApi {
// 	return &coreClient{client: client}
// }

// func PostFile[TReq api.Request, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.Method) (*TResp, error) {

// 	c, ok := s.(*APIClient)

// 	if !ok {
// 		return nil, errors.New("invalid client")
// 	}
// 	buf := new(bytes.Buffer)

// 	// Prepare a form that you will submit to that URL.
// 	if w, fs, err := form.Marshal(buf, method, r); err != nil {
// 		w.Close()
// 		return nil, err
// 	} else {
// 		defer w.Close()

// 		// Only set a timeout if one isn't already set
// 		var cancel context.CancelFunc
// 		if _, ok := ctx.Deadline(); !ok {
// 			ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
// 			defer cancel()
// 		}

// 		u := c.BaseURL
// 		req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, u.String(), buf)
// 		if err != nil {
// 			return nil, err
// 		}

// 		req.Header.Add("Content-Length", fmt.Sprintf("%d", fs))
// 		req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", w.Boundary()))

// 		return Do[TResp](c.httpClient, req)
// 	}
// }

// func List[TResp api.Response](c SynologyClient, ctx context.Context, method api.Method) (*TResp, error) {

// 	return Get[api.Request, TResp](c, ctx, nil, method)
// }

// func GetEncode[TReq api.EncodeRequest, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.Method) (*TResp, error) {
// 	return Get[TReq, TResp](s, ctx, r, method)
// }

// func Post[TReq api.Request, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.Method) (*TResp, error) {
// 	c, ok := s.(*APIClient)
// 	if !ok {
// 		return nil, errors.New("invalid client")
// 	}

// 	qu, err := util.Query(method, r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	u := c.BaseURL

// 	resp, err := c.httpClient.PostForm(u.String(), qu)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return handleResponse[TResp](resp)
// }

// func Get[TReq api.Request, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.Method) (*TResp, error) {
// 	c, ok := s.(*APIClient)
// 	if !ok {
// 		return nil, errors.New("invalid client")
// 	}

// 	aq, err := query.Values(method) //.AsApiParams())
// 	if err != nil {
// 		return nil, err
// 	}
// 	dq, err := query.Values(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	qu := maps.Clone(c.BaseURL.Query())
// 	maps.Copy(qu, aq)
// 	maps.Copy(qu, dq)

// 	u := c.BaseURL

// 	u.RawQuery = qu.Encode()

// 	// Only set a timeout if one isn't already set
// 	var cancel context.CancelFunc
// 	if _, ok := ctx.Deadline(); !ok {
// 		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
// 		defer cancel()
// 	}

// 	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return Do[TResp](c.httpClient, req)
// }

// func download(r io.ReadCloser) (interface{}, error) {
// 	var buf bytes.Buffer
// 	_, err := io.Copy(&buf, r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &form.File{
// 		Content: buf.String(),
// 		Name:    "download",
// 	}, nil
// }

// func Do[TResponse api.Response](client *retryablehttp.Client, req *retryablehttp.Request) (*TResponse, error) {
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		_, _ = io.ReadAll(resp.Body)
// 		_ = resp.Body.Close()
// 	}()

// 	return handleResponse[TResponse](resp)
// }

// func handleResponse[TResponse api.Response](resp *http.Response) (*TResponse, error) {
// 	var synoResponse api.ApiResponse[TResponse]

// 	contentType := resp.Header.Get("Content-Type")
// 	contentType = strings.Split(contentType, ";")[0]

// 	switch contentType {
// 	case "application/json":
// 		if err := json.NewDecoder(resp.Body).Decode(&synoResponse); err != nil {
// 			return nil, err
// 		}
// 	case "application/octet-stream":
// 		resp, err := download(resp.Body)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if resp, ok := resp.(*TResponse); ok {
// 			return resp, nil
// 		} else {
// 			return nil, errors.New("invalid response")
// 		}
// 	}

// 	if synoResponse.Success {
// 		return &synoResponse.Data, nil
// 	} else {
// 		return nil, handleErrors(synoResponse, api.GlobalErrors)
// 	}
// }

// func handleErrors[T any | api.ApiError](response api.ApiResponse[T], knownErrors api.ErrorSummary) error {
// 	if response.Error.Code == 0 {
// 		return nil
// 	}

// 	var result error

// 	if errDesc, ok := knownErrors[response.Error.Code]; ok {
// 		result = multierror.Append(result, fmt.Errorf("api response error code %d: %v", response.Error.Code, errDesc))
// 	} else {
// 		result = multierror.Append(result, fmt.Errorf("api response error code %d: %v", response.Error.Code, response.Error))
// 	}

// 	if response.Error.Errors != nil {
// 		for i, err := range response.Error.Errors {
// 			if errDesc, ok := knownErrors[err.Code]; ok {
// 				result = multierror.Append(result, fmt.Errorf("api response error[%d] code %d: %v", i, err.Code, errDesc))
// 			} else {
// 				result = multierror.Append(result, fmt.Errorf("api response error[%d] code %d: %v", i, err.Code, err))
// 			}
// 		}
// 	}

// 	return result
// }
