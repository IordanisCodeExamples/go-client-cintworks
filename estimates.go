package client

import (
	"encoding/json"
	"fmt"
)

const (
	estimatesEndpoint = "/ordering/FeasibilityEstimates?version=2"
)

// GetFeasibilityEstimates sends a survey settings request to the Cint Demand API
// to obtain feasibility estimates. It takes a SurveySettingsRequest struct as input,
// serializes it to JSON, sends it as a payload in a POST request, and then deserializes
// the received response into a SurveyResponse struct.
func (c *Client) GetFeasibilityEstimates(surveySettings SurveySettingsRequest) (SurveyResponse, error) {
	payloadBytes, err := json.Marshal(surveySettings)
	if err != nil {
		return SurveyResponse{}, fmt.Errorf("failed to marshal survey settings: %w", err)
	}

	responseBody, err := c.post(estimatesEndpoint, string(payloadBytes))
	if err != nil {
		return SurveyResponse{}, fmt.Errorf("failed to post to feasibility estimates endpoint: %w", err)
	}

	var response SurveyResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return SurveyResponse{}, fmt.Errorf("failed to unmarshal survey response: %w", err)
	}

	return response, nil

}
