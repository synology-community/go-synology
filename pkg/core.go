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

		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		u := c.BaseURL
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), buf)
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

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
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
		File: form.File{
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

	if synoResponse.Success {
		return &synoResponse.Data, nil
	} else {
		return nil, handleErrors(synoResponse, api.GlobalErrors)
	}

	// response.SetError(handleErrors(synoResponse, response, api.GlobalErrors))
	// return nil
}

func handleErrors[T any | api.ApiError](response api.ApiResponse[T], knownErrors api.ErrorSummary) error {
	if response.Error.Code == 0 {
		return nil
	}

	if errDesc, ok := knownErrors[response.Error.Code]; ok {
		return fmt.Errorf("Api Error: %v", errDesc)
	} else {
		return fmt.Errorf("Api Error: %v", response.Error)
	}
}
