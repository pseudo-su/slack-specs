package conversations

import (
	// "fmt"
	"context"

	"slack-specs-test-harness-go/pkg"

	"github.com/stretchr/testify/assert"
	// s "slack-specs-test-harness-go/smoke/suite"
)

func (suite *TestSuite) TestListConversations() {
	t := suite.T()
	ctx := context.Background()
	params := pkg.ConversationsListParams{}
	r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// fmt.Println(s.PrettyJSON(t, r.Body))
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
}

func (suite *TestSuite) TestListConversationsExcludeArchived() {
	t := suite.T()
	t.Skip()
	// ctx := context.Background()
	// params := pkg.ConversationsListParams{}
	// r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// // fmt.Println(s.PrettyJSON(t, r.Body))
	// assert.Equal(t, nil, err)
	// assert.NotNil(t, r.JSON200)

	t.FailNow()
}

func (suite *TestSuite) TestListConversationsPaginated() {
	t := suite.T()
	t.Skip()
	// ctx := context.Background()
	// params := pkg.ConversationsListParams{}
	// r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// // fmt.Println(s.PrettyJSON(t, r.Body))
	// assert.Equal(t, nil, err)
	// assert.NotNil(t, r.JSON200)

	t.FailNow()
}

func (suite *TestSuite) TestListConversationsTypeChannel() {
	t := suite.T()
	t.Skip()
	// ctx := context.Background()
	// params := pkg.ConversationsListParams{}
	// r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// // fmt.Println(s.PrettyJSON(t, r.Body))
	// assert.Equal(t, nil, err)
	// assert.NotNil(t, r.JSON200)

	t.FailNow()
}

func (suite *TestSuite) TestListConversationsTypeGroup() {
	t := suite.T()
	t.Skip()
	// ctx := context.Background()
	// params := pkg.ConversationsListParams{}
	// r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// // fmt.Println(s.PrettyJSON(t, r.Body))
	// assert.Equal(t, nil, err)
	// assert.NotNil(t, r.JSON200)

	t.FailNow()
}

func (suite *TestSuite) TestListConversationsTypeIM() {
	t := suite.T()
	t.Skip()
	// ctx := context.Background()
	// params := pkg.ConversationsListParams{}
	// r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// // fmt.Println(s.PrettyJSON(t, r.Body))
	// assert.Equal(t, nil, err)
	// assert.NotNil(t, r.JSON200)

	t.FailNow()
}

func (suite *TestSuite) TestListConversationsTypeMPIM() {
	t := suite.T()
	t.Skip()
	// ctx := context.Background()
	// params := pkg.ConversationsListParams{}
	// r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// // fmt.Println(s.PrettyJSON(t, r.Body))
	// assert.Equal(t, nil, err)
	// assert.NotNil(t, r.JSON200)

	t.FailNow()
}
