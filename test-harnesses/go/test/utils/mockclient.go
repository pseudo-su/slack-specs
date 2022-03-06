package utils

import (
	"net/http"
	"slack-specs-test-harness-go/pkg"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type HTTPClientMock struct {
	mock.Mock
}

func (c *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)

	arg1, arg2 := args.Get(0), args.Error(1)
	if arg1 == nil {
		return nil, arg2
	}
	return arg1.(*http.Response), arg2
}

func NewMockedClientWithResponses(t *testing.T) (*HTTPClientMock, pkg.ClientWithResponsesInterface) {
	mockClient := new(HTTPClientMock)

	mockClientWithValidation := NewHttpClientWithValidation(t, "../../../../../dist/openapi.manicured.yaml", mockClient)

	apiClient, err := pkg.NewClientWithResponses("https://slack.com/api", pkg.WithHTTPClient(mockClientWithValidation))
	require.NoError(t, err)
	return mockClient, apiClient
}
