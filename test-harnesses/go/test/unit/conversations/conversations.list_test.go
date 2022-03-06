package conversations

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"slack-specs-test-harness-go/pkg"
	"slack-specs-test-harness-go/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListConversationsCommonSuccess(t *testing.T) {
	httpClientMock, apiClient := utils.NewMockedClientWithResponses(t)

	// setup mock
	responseFixture := utils.LoadFixture(t, "conversations.list", "ResponseBody", "common")
	body := io.NopCloser(bytes.NewReader(responseFixture))
	httpClientMock.On("Do", mock.Anything).Return(&http.Response{
		Status: "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: body,
	}, nil)

	// make request
	ctx := context.Background()
	params := pkg.ConversationsListParams{}
	r, err := apiClient.ConversationsListWithResponse(ctx, &params, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer xoxb-XXXXXXXXXXXX-XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXX")
		return nil
	})

	// assert response
	require.NoError(t, err)
	assert.NotNil(t, r.JSON200)
	assert.Nil(t, r.JSONDefault)
}
func TestListConversationsAlternateSuccess(t *testing.T) {
	httpClientMock, apiClient := utils.NewMockedClientWithResponses(t)
	// setup mock
	responseFixture := utils.LoadFixture(t, "conversations.list", "ResponseBody", "alternate")
	httpClientMock.On("Do", mock.Anything).Return(&http.Response{
		Status: "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(responseFixture)),
	}, nil)

	// make request
	ctx := context.Background()
	params := pkg.ConversationsListParams{}
	r, err := apiClient.ConversationsListWithResponse(ctx, &params, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer xoxb-XXXXXXXXXXXX-XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXX")
		return nil
	})

	// assert response
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
	assert.Nil(t, r.JSONDefault)
}

func TestListConversationsCommonError(t *testing.T) {
	httpClientMock, apiClient := utils.NewMockedClientWithResponses(t)

	// setup mock
	responseFixture := utils.LoadFixture(t, "conversations.list", "ErrorResponseBody", "common")
	httpClientMock.On("Do", mock.Anything).Return(&http.Response{
		Status: "200 OK",
		StatusCode: 403,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(responseFixture)),
	}, nil)

	// make request
	ctx := context.Background()
	params := pkg.ConversationsListParams{}
	r, err := apiClient.ConversationsListWithResponse(ctx, &params, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer xoxb-XXXXXXXXXXXX-XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXX")
		return nil
	})

	// assert response
	assert.Equal(t, nil, err)
	assert.Nil(t, r.JSON200)
	assert.NotNil(t, r.JSONDefault)
}
