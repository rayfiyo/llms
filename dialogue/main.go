package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/api"
	"github.com/rayfiyo/llms/dialogue/cmd"
	"github.com/rayfiyo/llms/dialogue/cmd/files"
	"github.com/rayfiyo/llms/dialogue/models"
)

func main() {
	mode := flag.String(
		"mode", "generate", "Mode to use: 'chat' or 'generate'.")
	model := flag.String(
		"model", "Llama-3-Swallow-70B-Instruct-v0.1-Q8_0",
		"Model name.")
	cyclesLimit := flag.Int("limit", 12, "Limit number of sends cycles.")
	head := flag.String("head", "", "Prompt head . Fixed statement.")
	flag.Parse()
	prompt := flag.Arg(0)

	fileName := cmd.GenerateFileName()

	client := api.NewClient("http://172.27.167.204:11434")

	var content string
	var err error

	for i := 1; i < *cyclesLimit+1; i++ {
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
			log.Fatalf(
				"Invalid mode: %s. Use 'chat' or 'generate'", *mode,
			)
		}

		if err != nil {
			log.Fatalf("Error@%d: %v", i, err)
		}

		// 出力系
		if err := files.Append(fileName, "## "+fmt.Sprint(i)); err != nil {
			log.Fatalf("Error appending to file 1@%d: %v", i, err)
		}
		if err := files.Append(fileName, content); err != nil {
			log.Fatalf("Error appending to file 2@%d: %v", i, err)
		}
		fmt.Println("- - - - - - - - - - - -")
		fmt.Printf("%3d:\n%s\n", i, content)

		// 後処理
		prompt = "出力は日本語で行ってください。" + *head + content
	}
}
