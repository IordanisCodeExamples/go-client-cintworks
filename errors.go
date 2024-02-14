package client

import "errors"

var (
	// ErrNoDomain is returned when the domain is missing from the client configuration.
	ErrNoDomain = errors.New(`error_no_domain`)
	// ErrNoAPIKey is returned when the API key is missing from the client configuration.
	ErrNoAPIKey = errors.New(`error_no_api_key`)
	// ErrNoHTTPClient is returned when the HTTP client is missing from the client configuration.
	ErrNoHTTPClient = errors.New(`error_no_http_client`)
)
