package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	url    string
	secret string
}

func NewClient(url, secret string) *Client {
	return &Client{url: url, secret: secret}
}

func (c *Client) Query(query string, variables map[string]interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", c.secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
