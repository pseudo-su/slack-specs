package smoke

import (
	"context"

	pkg "github.com/pseudo-su/slack-specs/pkg/complete"

	"github.com/stretchr/testify/assert"
)

func (suite *TestSuite) TestListConversations() {
	t := suite.T()
	ctx := context.Background()
	params := pkg.ConversationsListParams{
		Token: &suite.suiteCtx.envValues.APIToken,
	}
	r, err := suite.apiClient.ConversationsListWithResponse(ctx, &params)

	// fmt.Println(prettyJSON(t, r.Body))
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
}
