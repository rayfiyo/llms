package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/api"
	"github.com/rayfiyo/llms/dialogue/models"
)

func main() {
	mode := flag.String("mode", "chat", "Mode to use: 'chat' or 'generate'")
	model := flag.String("model", "Llama-3-Swallow-70B-Instruct-v0.1-Q8_0", "model name flag")
	flag.Parse()
	prompt := flag.Arg(0)

	client := api.NewClient("http://172.27.167.204:11434")

	var content string
	var err error

	switch *mode {
	case "chat":
		request := &models.ChatRequest{
			Model: *model,
			Messages: []models.Message{
				{Role: "user", Content: prompt},
			},
		}
		content, err = client.Chat(request)
	case "generate":
		request := &models.GenerateRequest{
			Model:  *model,
			Prompt: prompt,
		}
		content, err = client.Generate(request)
	default:
		log.Fatalf("Invalid mode: %s. Use 'chat' or 'generate'", *mode)
	}

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(content)

	/*
		for i := 1; ; i++ {
			fmt.Println("- - - - - -")
			fmt.Println(i)
			fmt.Println(answer)
			fmt.Println("- - - - - -")

			answer = ""

			requestBody = GenerateRequest{
				Model:  *model,
				Prompt: answer,
			}

			jsonBody, err = json.Marshal(requestBody)
			if err != nil {
				log.Fatal(err, i)
			}

			responses, err = client.Chat(jsonBody)
			if err != nil {
				log.Fatal(err, i)
			}

			var answer string
			for _, response := range responses {
				answer += response.Content
			}
		}
	*/
}
