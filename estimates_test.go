package client_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/junkd0g/go-client-cintworks"
	"github.com/stretchr/testify/assert"
)

func TestGetFeasibilityEstimatesSuccess(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a mock JSON response
		io.WriteString(w, `{"incidenceRate": 10, "lengthOfInterview": 20}`)
	}))
	defer server.Close()

	// Create a new client with the mock server's URL
	cl, err := client.New(server.URL, "test-api-key", server.Client())
	assert.NoError(t, err)

	// Call the method under test
	resp, err := cl.GetFeasibilityEstimates(client.SurveySettingsRequest{})
	assert.NoError(t, err)
	assert.Equal(t, 10, resp.IncidenceRate)
	assert.Equal(t, 20, resp.LengthOfInterview)
}

func TestGetFeasibilityEstimatesUnMarshalingError(t *testing.T) {
	cl, err := client.New("http://example.com", "test-api-key", &http.Client{})
	assert.NoError(t, err)
	_, err = cl.GetFeasibilityEstimates(client.SurveySettingsRequest{})
	assert.Contains(t, err.Error(), "failed to unmarshal survey response")
}
