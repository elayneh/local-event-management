package utils

import (
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)

// headersTransport is a custom HTTP transport that adds headers to each request
type headersTransport struct {
	headers http.Header
	base    http.RoundTripper
}

func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, values := range t.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return t.base.RoundTrip(req)
}

var client *graphql.Client

func Client() *graphql.Client {
	if client == nil {
		// Get the Hasura URL from the environment variable
		hasuraURL := os.Getenv("HASURA_GRAPHQL_URL")
		if hasuraURL == "" {
			hasuraURL = "http://minabtest-hasura-1:8080/v1/graphql" // Default value if env variable is not set
		}

		// Set up the HTTP client
		httpClient := &http.Client{}

		// Set up the GraphQL client with the custom HTTP client
		client = graphql.NewClient(hasuraURL, graphql.WithHTTPClient(httpClient))
	}

	return client
}
