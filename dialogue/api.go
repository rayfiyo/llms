package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *APIClient) Chat(model string, messages []Message) ([]Message, error) {
	requestBody := ChatRequest{
		Model:    model,
		Messages: messages,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responses []Message
	decoder := json.NewDecoder(resp.Body)
	for {
		var chatResp ChatResponse
		if err := decoder.Decode(&chatResp); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		responses = append(responses, chatResp.Message)
	}

	return responses, nil
}
