// Package client provides a specialized HTTP client for interacting with the Cint Demand API.
// It facilitates the construction and execution of API requests, handling authentication,
// and parsing responses.
package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client encapsulates the necessary configuration for making authenticated requests to the Cint Demand API.
// It holds the API's base domain, an API key for authentication, and an http.Client to execute the requests.
type Client struct {
	Domain string
	ApiKey string
	Client *http.Client
}

// New initializes a new Client with the specified domain, API key, and an http.Client.
// It validates the inputs to ensure that the client is configured with all necessary information
// for making requests to the Cint Demand API. Returns an error if any required field is missing.
func New(domain, apiKey string, httpClient *http.Client) (*Client, error) {
	switch {
	case domain == ``:
		return nil, ErrNoDomain
	case apiKey == ``:
		return nil, ErrNoAPIKey
	case httpClient == nil:
		return nil, ErrNoHTTPClient
	}

	return &Client{
		Domain: domain,
		ApiKey: apiKey,
		Client: httpClient,
	}, nil
}

// post constructs and sends a POST request to a specified endpoint of the Cint Demand API with the given payload.
// It automatically adds necessary headers for content type and authentication using the API key.
// The method constructs the full request URL from the client's base domain and endpoint,
// ensuring proper URL formatting. Returns the response body as bytes or an error if the request
// fails or cannot be sent.
func (c *Client) post(endpoint string, payload string) ([]byte, error) {
	baseURL, err := url.Parse(c.Domain)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}
	fullURL := baseURL.ResolveReference(endpointURL).String()

	req, err := http.NewRequest(http.MethodPost, fullURL, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", acceptHeader)
	req.Header.Add("content-type", contentTypeHeader)
	req.Header.Add("X-Api-Key", c.ApiKey)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
