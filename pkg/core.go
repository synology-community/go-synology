package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/util/form"
)

var defaultTimeout = 15 * time.Second

func LookupMethod(callingFunc string) (*api.APIMethod, error) {
	var api, method string

	sp := strings.Split(callingFunc, ".")

	if len(sp) < 2 {
		return nil, errors.New("invalid caller function")
	}

	api = sp[0]
	method = sp[1]

	if res, ok := MethodLookup[api][method]; ok {
		return &res, nil
	} else {
		return nil, errors.New("method not found")
	}
}

func Post[TReq api.Request, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.APIMethod) (*TResp, error) {

	c, ok := s.(*APIClient)

	if !ok {
		return nil, errors.New("invalid client")
	}
	buf := new(bytes.Buffer)

	//buf.WriteString(fmt.Sprintf("multipart/form-data; boundary=%s\n\n\\--AaB03x", "--AaB03x"))

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

		u := c.BaseURL
		req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, u.String(), buf)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Length", fmt.Sprintf("%d", fs))
		req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", w.Boundary()))

		return Do[TResp](c.httpClient, req)
	}
}

func List[TResp api.Response](c SynologyClient, ctx context.Context, method api.APIMethod) (*TResp, error) {

	return Get[api.Request, TResp](c, ctx, nil, method)
}

func GetEncode[TReq api.EncodeRequest, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.APIMethod) (*TResp, error) {
	return Get[TReq, TResp](s, ctx, r, method)
}

func Get[TReq api.Request, TResp api.Response](s SynologyClient, ctx context.Context, r *TReq, method api.APIMethod) (*TResp, error) {
	c, ok := s.(*APIClient)
	if !ok {
		return nil, errors.New("invalid client")
	}

	aq, err := query.Values(method) //.AsApiParams())
	if err != nil {
		return nil, err
	}
	dq, err := query.Values(r)
	if err != nil {
		return nil, err
	}

	qu := maps.Clone(c.BaseURL.Query())
	maps.Copy(qu, aq)
	maps.Copy(qu, dq)

	u := c.BaseURL

	// if len(params) > 0 {
	// 	u = params[0]
	// } else {
	// 	u = c.BaseURL
	// }

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

	return Do[TResp](c.httpClient, req)
}

func download(r io.ReadCloser) (interface{}, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	dlr := filestation.DownloadResponse{
		File: form.File{
			Content: buf.String(),
			Name:    "download",
		},
	}

	return &dlr, nil
}

func Do[TResponse api.Response](client *retryablehttp.Client, req *retryablehttp.Request) (*TResponse, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}()

	// application/octet-stream

	if resp.Header.Get("Content-Type") == "application/octet-stream" {
		resp, err := download(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp, ok := resp.(*TResponse); ok {
			return resp, nil
		} else {
			return nil, errors.New("invalid response")
		}
	}
	var synoResponse api.ApiResponse[TResponse]

	if err := json.NewDecoder(resp.Body).Decode(&synoResponse); err != nil {
		return nil, err
	}

	if synoResponse.Success {
		return &synoResponse.Data, nil
	} else {
		return nil, handleErrors(synoResponse, api.GlobalErrors)
	}

	// response.SetError(handleErrors(synoResponse, response, api.GlobalErrors))
	// return nil
}

// func Retry(ctx context.Context, delay time.Duration, fn func() error) error {
// 	deadline, ok := ctx.Deadline()
// 	if !ok {
// 		deadline = time.Now().Add(60 * time.Second)
// 	}

// 	for {
// 		err := fn()
// 		if err == nil {
// 			return nil
// 		}

// 		delay := 5 * time.Second
// 		for {
// 			if err := fn(); err != nil {
// 				return nil
// 			}
// 			if task.Finished {
// 				return task, nil
// 			}
// 			if time.Now().After(deadline.Add(delay)) {
// 				return nil, fmt.Errorf("Timeout waiting for task to complete")
// 			}
// 			time.Sleep(delay)
// 		}
// 	}
// }

func handleErrors[T any | api.ApiError](response api.ApiResponse[T], knownErrors api.ErrorSummary) error {
	if response.Error.Code == 0 {
		return nil
	}

	var result error

	if errDesc, ok := knownErrors[response.Error.Code]; ok {
		result = multierror.Append(result, fmt.Errorf("api response error: response code %d => %v", response.Error.Code, errDesc))
	} else {
		result = multierror.Append(result, fmt.Errorf("api response error: response code %d => %v", response.Error.Code, response.Error))
	}

	if response.Error.Errors != nil {
		for i, err := range response.Error.Errors {
			if errDesc, ok := knownErrors[err.Code]; ok {
				result = multierror.Append(result, fmt.Errorf("api response error[%d]: response code %d => %v", i, err.Code, errDesc))
			} else {
				result = multierror.Append(result, fmt.Errorf("api response error[%d]: response code %d => %v", i, err.Code, err))
			}
		}
	}

	return result
}
