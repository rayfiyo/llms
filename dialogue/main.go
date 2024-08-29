package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/api"
	"github.com/rayfiyo/llms/dialogue/models"
)

func main() {
	var (
		url = flag.String("url", "http://172.27.167.204:11434/api/generate", "api client URL flag")
		// url = flag.String("url", "http://172.27.167.204:11434/api/chat", "api client URL flag")
		model = flag.String("model", "Llama-3-Swallow-70B-Instruct-v0.1-Q8_0", "model name flag")
	)

	flag.Parse()
	prompt := flag.Arg(0)

	client := api.NewClient(*url)

	/*
		requestBody := GenerateRequest{
			Model:  *model,
			Prompt: prompt,
		}
	*/

	request := &models.ChatRequest{
		Model: *model,
		Messages: []models.Message{
			{Role: "user", Content: prompt},
		},
	}

	responses, err := client.Chat(request)
	if err != nil {
		log.Fatalf("Error during chat: %v %d", err, 0)
	}

	var answer string
	for _, resp := range responses {
		fmt.Print(resp.Message.Content)
		answer += resp.Message.Content
	}

	fmt.Println(answer)

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
