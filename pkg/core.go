package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/util/form"
)

type callerKey string

const callerContextKey callerKey = "caller"

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

func Post[TReq api.Request, TResp api.Response](s SynologyClient, r *TReq, method api.APIMethod) (*TResp, error) {

	c, ok := s.(*APIClient)

	if !ok {
		return nil, errors.New("invalid client")
	}

	req := api.NewRequest(method.AsApiParams(), r)

	// Prepare a form that you will submit to that URL.
	if b, err := form.Marshal(req); err != nil {
		return nil, err
	} else {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		u := c.BaseURL
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(b))
		if err != nil {
			return nil, err
		}

		return Do[TResp](c.httpClient, req)
	}
}

func List[TResp api.Response](c SynologyClient, method api.APIMethod) (*TResp, error) {

	return Get[api.Request, TResp](c, nil, method)
}

func GetEncode[TReq api.EncodeRequest, TResp api.Response](s SynologyClient, r *TReq, method api.APIMethod) (*TResp, error) {
	return Get[TReq, TResp](s, r, method)
}

func Get[TReq api.Request, TResp api.Response](s SynologyClient, r *TReq, method api.APIMethod) (*TResp, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
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
		File: &form.File{
			Content: buf.String(),
			Name:    "download",
		},
	}

	return &dlr, nil
}

func Do[TResponse api.Response](client *http.Client, req *http.Request) (*TResponse, error) {
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

	return &synoResponse.Data, nil

	// response.SetError(handleErrors(synoResponse, response, api.GlobalErrors))
	// return nil
}

func handleErrors[T any | api.ApiError](response api.ApiResponse[T], errorDescriber api.ErrorDescriber, knownErrors api.ErrorSummary) api.ApiError {
	err := api.ApiError{
		Code: response.Error.Code,
	}
	if response.Error.Code == 0 {
		return err
	}

	combinedKnownErrors := append(errorDescriber.ErrorSummaries(), knownErrors)
	err.Summary = api.DescribeError(err.Code, combinedKnownErrors...)
	for _, e := range response.Error.Errors {
		item := api.ErrorItem{
			Code:    e.Code,
			Summary: api.DescribeError(e.Code, combinedKnownErrors...),
		}
		if len(e.Details) > 0 {
			item.Details = make(api.ErrorFields)
			for k, v := range e.Details {
				item.Details[k] = v
			}
		}
		err.Errors = append(err.Errors, item)
	}

	return err
}
