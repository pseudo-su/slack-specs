package conversations

import (
	"testing"


	testifySuite "github.com/stretchr/testify/suite"
	suite "slack-specs-test-harness-go/smoke/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRunSuite(t *testing.T) {
	testifySuite.Run(t, new(TestSuite))
}
