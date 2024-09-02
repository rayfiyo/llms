package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func (c *Client) Chat(req *models.ChatRequest) (string, error) {
	return c.sendRequest("/api/chat", req, "chat")
}

func (c *Client) Generate(req *models.GenerateRequest) (string, error) {
	return c.sendRequest("/api/generate", req, "generate")
}

func (c *Client) sendRequest(endpoint string, req interface{}, mode string) (string, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	resp, err := c.HTTPClient.Post(c.BaseURL+endpoint,
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d\n%s",
			resp.StatusCode, resp.Body)
	}

	var content strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		switch mode {
		case "chat":
			var response models.ChatResponse
			if err := json.Unmarshal(
				scanner.Bytes(), &response,
			); err != nil {
				return "", fmt.Errorf("Error unmarshaling chat response: %v", err)
			}
			content.WriteString(response.Message.Content)
		case "generate":
			var response models.GenerateResponse
			if err := json.Unmarshal(
				scanner.Bytes(), &response,
			); err != nil {
				return "", fmt.Errorf("Error unmarshaling chat response: %v", err)
			}
			content.WriteString(response.Response)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	return content.String(), nil
}
