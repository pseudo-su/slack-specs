package utils

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)

	arg1, arg2 := args.Get(0), args.Error(1)
	if arg1 == nil {
		return nil, arg2
	}
	return arg1.(*http.Response), arg2
}
