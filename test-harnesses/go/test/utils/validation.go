package utils

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/stretchr/testify/require"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPClientWithValidation struct {
	t              *testing.T
	origHTTPClient HTTPDoer
	spec           *openapi3.T
	router         routers.Router
}

func (c *HTTPClientWithValidation) Do(req *http.Request) (*http.Response, error) {
	// Find route
	route, pathParams, err := c.router.FindRoute(req)
	require.NoError(c.t, err)

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
		// TODO: are these required?
		// QueryParams: nil,
		// ParamDecoder: nil,
		Options: &openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false,
			IncludeResponseStatus: true,
			MultiError:            true,
			AuthenticationFunc: func(c context.Context, ai *openapi3filter.AuthenticationInput) error {
				fmt.Println(ai)
				return nil
			},
		},
	}
	err = openapi3filter.ValidateRequest(req.Context(), requestValidationInput)
	require.NoError(c.t, err, "Request validation failed")

	res, err := c.origHTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	// read the response body to a variable
    bodyBytes, _ := ioutil.ReadAll(res.Body)
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

    //reset the response body to the original unread state
    defer func() {
		res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}()


	// Validate response
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 res.StatusCode,
		Header:                 res.Header,
		Body:                   res.Body,
		Options: &openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false,
			IncludeResponseStatus: true,
			MultiError:            true,
			AuthenticationFunc: func(c context.Context, ai *openapi3filter.AuthenticationInput) error {
				fmt.Println(ai)
				return nil
			},
		},
	}

	// Validate response.
	err = openapi3filter.ValidateResponse(req.Context(), responseValidationInput)
	require.NoError(c.t, err, "Response validation failed")

	return res, nil
}

func NewHttpClientWithValidation(t *testing.T, filepath string, httpClient HTTPDoer) HTTPDoer {
	loader := openapi3.NewLoader()

	spec, err := loader.LoadFromFile(filepath)
	require.NoError(t, err)
	err = spec.Validate(loader.Context)
	require.NoError(t, err)

	router, err := gorillamux.NewRouter(spec)
	require.NoError(t, err)

	client := &HTTPClientWithValidation{
		t:              t,
		origHTTPClient: httpClient,
		spec:           spec,
		router:         router,
	}

	return client
}
