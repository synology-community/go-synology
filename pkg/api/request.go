package api

import (
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/google/go-querystring/query"
	"github.com/synology-community/go-synology/pkg/util"
)

type Request interface {
	any
}

type EncodeRequest interface {
	Request
	query.Encoder
}

type ApiParams struct {
	Version int    `form:"version" query:"version"`
	API     string `form:"api" query:"version"`
	Method  string `form:"method" query:"version"`
}

type ApiRequest[TRequest Request] struct {
	ApiParams
	Request
}

type ApiRequestBuilder[TRequest Request] struct {
	baseRequest  ApiParams
	paramRequest TRequest
}

func (b *ApiRequestBuilder[TRequest]) WithVersion(version int) *ApiRequestBuilder[TRequest] {
	b.baseRequest.Version = version
	return b
}

func (b *ApiRequestBuilder[TRequest]) WithAPIName(api string) *ApiRequestBuilder[TRequest] {
	b.baseRequest.API = api
	return b
}

func (b *ApiRequestBuilder[TRequest]) WithAPIMethod(method string) *ApiRequestBuilder[TRequest] {
	b.baseRequest.Method = method
	return b
}

func (b *ApiRequestBuilder[TRequest]) WithRequest(request TRequest) *ApiRequestBuilder[TRequest] {
	b.paramRequest = request
	return b
}

func (b *ApiRequestBuilder[TRequest]) Build() ApiRequest[TRequest] {
	return ApiRequest[TRequest]{
		ApiParams: b.baseRequest,
		Request:   b.paramRequest,
	}
}

func (b *ApiRequestBuilder[TRequest]) With(callback func(base *ApiParams, params *TRequest)) *ApiRequestBuilder[TRequest] {
	callback(&b.baseRequest, &b.paramRequest)
	return b
}

func NewRequest[TRequest Request](apiParams ApiParams, request TRequest) ApiRequest[TRequest] {
	return ApiRequest[TRequest]{
		ApiParams: apiParams,
		Request:   request,
	}
}

func (l ApiRequest[Request]) EncodeValues(_ string, _ *url.Values) error {
	return nil
}

func RequestBuilder[TRequest Request](params TRequest) *ApiRequestBuilder[TRequest] {
	b := &ApiRequestBuilder[TRequest]{
		baseRequest:  ApiParams{},
		paramRequest: params,
	}

	caller, err := util.GetCaller()
	if err != nil {
		log.Fatalf("[ERROR] Failed to get caller: %v", err)
		return b
	}

	cParts := strings.Split(caller, ".")
	cName := cParts[len(cParts)-1]

	log.Infof("[INFO] called from %s\n", cName)

	lookup := map[string]struct {
		Name    string
		Method  string
		Version int
	}{
		"Login": {
			Name:    "SYNO.API.Auth",
			Method:  "login",
			Version: 7,
		},
	}

	if a, ok := lookup[cName]; ok {
		b.baseRequest.Version = a.Version
		b.baseRequest.API = a.Name
		b.baseRequest.Method = a.Method
	}

	return b
}

func CreateRequest[TRequest Request](version int, apiName string, apiMethod string, request TRequest) ApiRequest[TRequest] {
	b := RequestBuilder(request)
	return b.WithVersion(version).WithAPIName(apiName).WithAPIMethod(apiMethod).Build()
}

func GetRequest[TRequest Request](v TRequest) (ApiRequest[TRequest], error) {
	b := RequestBuilder(v)
	return b.Build(), nil
}
