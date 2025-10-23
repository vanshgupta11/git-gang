package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client handles communication with LeetCode API
type Client struct {
	client  *http.Client
	headers map[string]string
}

// New creates a new LeetCode API client
func New(session, csrf string) *Client {
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Origin":        "https://leetcode.com",
		"User-Agent":    "Mozilla/5.0",
		"x-csrftoken":   csrf,
		"Cookie":        fmt.Sprintf("LEETCODE_SESSION=%s; csrftoken=%s;", session, csrf),
	}

	return &Client{
		client:  &http.Client{},
		headers: headers,
	}
}

// GraphQLPayload represents a GraphQL query payload
type GraphQLPayload struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

// FetchGraphQL executes a GraphQL query against LeetCode API
func (c *Client) FetchGraphQL(payload *GraphQLPayload) (map[string]interface{}, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// Add headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}