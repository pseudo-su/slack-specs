package utils

import "slack-specs-test-harness-go/pkg"

func NewClientWithResponses(serverBaseURL string) (pkg.ClientWithResponsesInterface, error) {

	apiClient, err := pkg.NewClientWithResponses(serverBaseURL)
	if err != nil {
		return nil, err
	}

	return apiClient, nil
}
