package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

func main() {
	var (
		url = flag.String("str", "http://172.27.167.204:11434/api/generate", "api client URL flag")
		// url = flag.String("str", "http://172.27.167.204:11434/api/chat", "api client URL flag")
		model = flag.String("str", "Llama-3-Swallow-70B-Instruct-v0.1-Q8_0", "model name flag")
	)

	flag.Parse()
	prompt := flag.Arg(0)

	client := NewAPIClient(*url)

	requestBody := GenerateRequest{
		Model:  *model,
		Prompt: prompt,
	}

	/*
		requestBody := ChatRequest{
			Model: *model,
			Messages: []Message{
				{Role: "user", Content: prompt},
			},
		}
	*/

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err, 0)
	}

	responses, err := client.Chat(jsonBody)
	if err != nil {
		log.Fatal(err, 0)
	}

	var answer string
	for _, response := range responses {
		answer += response.Content
	}

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
}
