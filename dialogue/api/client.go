package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rayfiyo/llms/dialogue/models"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) Chat(req *models.ChatRequest) ([]models.Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	resp, err := c.HTTPClient.Post(c.BaseURL+"/api/chat", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var responses []models.ChatResponse
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var chatResp models.ChatResponse
		if err := json.Unmarshal(scanner.Bytes(), &chatResp); err != nil {
			return nil, fmt.Errorf("error unmarshaling response: %w", err)
		}
		responses = append(responses, chatResp)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return responses, nil
}
