package suite

import (
	"fmt"

	"slack-specs-test-harness-go/pkg"
	"slack-specs-test-harness-go/test/utils"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Context  *SuiteContext
	ApiClient pkg.ClientWithResponsesInterface
}

//nolint:stylecheck
func (s *Suite) SetupSuite() {
	s.Context = ParseSuiteConfig()
	serverBaseURL := buildBaseURL(s.Context)

	apiClient, err := utils.NewClientWithResponses(serverBaseURL)
	if err != nil {
		require.NoError(s.T(), err)
	}

	s.ApiClient = apiClient
}

func (s *Suite) TeardownSuite() {}

func buildBaseURL(suiteCtx *SuiteContext) string {
	basepath := suiteCtx.ServiceBasepath()

	return fmt.Sprintf(
		"%s://%s:%v%s",
		suiteCtx.EnvValues.APIScheme,
		suiteCtx.EnvValues.APIHost,
		suiteCtx.EnvValues.APIPort,
		basepath,
	)
}

type SuiteContext struct {
	EnvValues *testSuiteEnv
}

func (suiteCtx *SuiteContext) Env() string {
	return suiteCtx.EnvValues.Env
}

func (suiteCtx *SuiteContext) ServerPort() int {
	return suiteCtx.EnvValues.APIPort
}

func (suiteCtx *SuiteContext) ServiceBasepath() string {
	return suiteCtx.EnvValues.APIBasepath
}

type testSuiteEnv struct {
	APIScheme   string `env:"API_SCHEME" envDefault:"https"`
	APIHost     string `env:"API_HOST" envDefault:"slack.com"`
	APIPort     int    `env:"API_PORT" envDefault:"443"`
	APIBasepath string `env:"API_BASEPATH" envDefault:"/api"`
	APIToken    string `env:"SLACK_OAUTH_ACCESS_TOKEN,required"`
	Env         string `env:"ENV" envDefault:"local"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"debug"`
}

func ParseSuiteConfig() *SuiteContext {
	envValues := &testSuiteEnv{}
	if err := env.Parse(envValues); err != nil {
		log.Fatalf("unable to find env var key: %v \n", err)
	}

	return &SuiteContext{
		EnvValues: envValues,
	}
}
