package main

import (
	"fmt"
	"log"
)

func main() {
	client := NewAPIClient("http://172.27.167.204:11434/api/chat")

	messages := []Message{
		{Role: "user", Content: "こんにちは！100文字程度で何か教えて下さい！"},
	}

	responses, err := client.Chat("llama3:70b", messages)
	if err != nil {
		log.Fatal(err)
	}

	for _, response := range responses {
		fmt.Print(response.Content)
	}
	fmt.Println()
}
