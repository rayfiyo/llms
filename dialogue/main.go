package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/api"
	"github.com/rayfiyo/llms/dialogue/models"
)

func main() {
	mode := flag.String("mode", "generate", "Mode to use: 'chat' or 'generate'")
	model := flag.String("model", "Llama-3-Swallow-70B-Instruct-v0.1-Q8_0", "model name flag")
	flag.Parse()
	prompt := flag.Arg(0)

	client := api.NewClient("http://172.27.167.204:11434")

	var content string
	var err error

	for i := 1; i < 10; i++ {
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
			log.Fatalf("Error@%d: %v", i, err)
		}

		fmt.Printf("- - - - - - - - - - - -")
		fmt.Printf("%3d:\n%s\n", i, content)
		prompt = content
	}
}
