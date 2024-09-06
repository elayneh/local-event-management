package hasura

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGraphQLRequest(query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	requestBody, err := json.Marshal(GraphQLRequest{
		Query:     query,
		Variables: variables,
	})

	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", graphqlEndpoint, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", "myadminsecretkey")
	req.Header.Set("x-hasura-role", "admin")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response GraphQLResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(response.Errors) > 0 {
		return &response, fmt.Errorf("graphql errors: %v", response.Errors)
	}
	return &response, nil
}
