package conversations

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"slack-specs-test-harness-go/pkg"
	"slack-specs-test-harness-go/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListConversationsCommonSuccess(t *testing.T) {
	mockClient := new(utils.MockClient)
	apiClient, err := pkg.NewClientWithResponses("https://fake-url", pkg.WithHTTPClient(mockClient))
	if err != nil {
		panic(err)
	}

	// setup mock
	responseFixture := utils.LoadFixture(t, "conversations.list", "ResponseBody", "common")
	mockClient.On("Do", mock.Anything).Return(&http.Response{
		Status: "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(responseFixture)),
	}, nil)

	// make request
	ctx := context.Background()
	token := "fake-token"
	params := pkg.ConversationsListParams{
		Token: &token,
	}
	r, err := apiClient.ConversationsListWithResponse(ctx, &params)

	// assert response
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
	assert.Nil(t, r.JSONDefault)
}
func TestListConversationsAlternateSuccess(t *testing.T) {
	mockClient := new(utils.MockClient)
	apiClient, err := pkg.NewClientWithResponses("https://fake-url", pkg.WithHTTPClient(mockClient))
	if err != nil {
		panic(err)
	}

	// setup mock
	responseFixture := utils.LoadFixture(t, "conversations.list", "ResponseBody", "alternate")
	mockClient.On("Do", mock.Anything).Return(&http.Response{
		Status: "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(responseFixture)),
	}, nil)

	// make request
	ctx := context.Background()
	token := "fake-token"
	params := pkg.ConversationsListParams{
		Token: &token,
	}
	r, err := apiClient.ConversationsListWithResponse(ctx, &params)

	// assert response
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
	assert.Nil(t, r.JSONDefault)
}

func TestListConversationsCommonError(t *testing.T) {
	mockClient := new(utils.MockClient)
	apiClient, err := pkg.NewClientWithResponses("https://fake-url", pkg.WithHTTPClient(mockClient))
	if err != nil {
		panic(err)
	}

	// setup mock
	responseFixture := utils.LoadFixture(t, "conversations.list", "ErrorResponseBody", "common")
	mockClient.On("Do", mock.Anything).Return(&http.Response{
		Status: "200 OK",
		StatusCode: 403,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(responseFixture)),
	}, nil)

	// make request
	ctx := context.Background()
	token := "fake-token"
	params := pkg.ConversationsListParams{
		Token: &token,
	}
	r, err := apiClient.ConversationsListWithResponse(ctx, &params)

	// assert response
	assert.Equal(t, nil, err)
	assert.Nil(t, r.JSON200)
	assert.NotNil(t, r.JSONDefault)
}
