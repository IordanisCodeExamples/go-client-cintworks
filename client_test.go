package client_test

import (
	"net/http"
	"testing"

	"github.com/junkd0g/go-client-cintworks"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {

	tests := []struct {
		name           string
		domain         string
		apiKey         string
		httpClient     *http.Client
		expectedClient *client.Client
		expectedError  error
	}{
		{
			name:       "Creates successfully a client object",
			domain:     "https://api.cint.com",
			apiKey:     "123456",
			httpClient: &http.Client{},
			expectedClient: &client.Client{
				Domain: "https://api.cint.com",
				ApiKey: "123456",
				Client: &http.Client{},
			},
			expectedError: nil,
		},
		{
			name:           "Fails to create a client object with no domain",
			domain:         "",
			apiKey:         "123456",
			httpClient:     &http.Client{},
			expectedClient: nil,
			expectedError:  client.ErrNoDomain,
		},
		{
			name:           "Fails to create a client object with no API key",
			domain:         "https://api.cint.com",
			apiKey:         "",
			httpClient:     &http.Client{},
			expectedClient: nil,
			expectedError:  client.ErrNoAPIKey,
		},
		{
			name:           "Fails to create a client object with no HTTP client",
			domain:         "https://api.cint.com",
			apiKey:         "123456",
			httpClient:     nil,
			expectedClient: nil,
			expectedError:  client.ErrNoHTTPClient,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cl, err := client.New(tc.domain, tc.apiKey, tc.httpClient)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedClient, cl)
		})
	}
}
