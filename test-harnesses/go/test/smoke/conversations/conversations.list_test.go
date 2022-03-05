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
	params := pkg.ConversationsListParams{
		Token: &suite.Context.EnvValues.APIToken,
	}
	r, err := suite.ApiClient.ConversationsListWithResponse(ctx, &params)

	// fmt.Println(s.PrettyJSON(t, r.Body))
	assert.Equal(t, nil, err)
	assert.NotNil(t, r.JSON200)
}
