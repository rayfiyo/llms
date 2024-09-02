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

func (c *Client) Chat(req *models.ChatRequest) (string, []int, error) {
	return c.sendRequest("/api/chat", req, "chat")
}

func (c *Client) Generate(req *models.GenerateRequest) (string, []int, error) {
	return c.sendRequest("/api/generate", req, "generate")
}

func (c *Client) sendRequest(endpoint string, req interface{}, mode string) (string, []int, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return "", nil, fmt.Errorf("error marshaling request: %w", err)
	}

	resp, err := c.HTTPClient.Post(c.BaseURL+endpoint,
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("unexpected status code: %d\n%s",
			resp.StatusCode, string(bufio.NewScanner(resp.Body).Bytes()),
		)
	}

	var content strings.Builder
	var context []int
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		switch mode {

		case "chat":
			var response models.ChatResponse
			if err := json.Unmarshal(
				scanner.Bytes(), &response,
			); err != nil {
				return "", nil, fmt.Errorf(
					"Error unmarshaling chat response: %v", err)
			}

			// 逐次標準出力
			fmt.Print(response.Message.Content)

			// 変数への書き込み
			content.WriteString(response.Message.Content)

		case "generate":
			var response models.GenerateResponse
			if err := json.Unmarshal(
				scanner.Bytes(), &response,
			); err != nil {
				return "", nil, fmt.Errorf(
					"Error unmarshaling chat response: %v", err)
			}

			// 逐次標準出力
			fmt.Print(response.Response)

			// 変数への書き込み
			content.WriteString(response.Response)
			context = response.Context
		}
	}
	fmt.Println("") // 文末調整用

	if err := scanner.Err(); err != nil {
		return "", context, fmt.Errorf("error reading response: %w", err)
	}

	return content.String(), context, nil
}
