package conversations

import (
	"context"

	pkg "slack-specs-test-harness-go/pkg"

	"github.com/stretchr/testify/assert"
)

func (suite *TestSuite) TestListConversations() {
	t := suite.T()
	ctx := context.Background()
	params := pkg.ConversationsListParams{
		Token: &suite.Context.EnvValues.APIToken,
	}
	r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// fmt.Println(prettyJSON(t, r.Body))
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
}
