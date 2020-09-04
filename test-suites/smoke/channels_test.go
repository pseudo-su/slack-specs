package smoke

import (
	"context"

	"github.com/pseudo-su/slack-specs/pkg"

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
	assert.Equal(t, pkg.DefsOkTrue(true), r.JSON200.Ok)
}
