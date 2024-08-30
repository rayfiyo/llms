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
	cyclesLimit := flag.Int("limit", 6, "Limit number of sends cycles.")
	head := flag.String("head", "", "Head of prompt. Fixed statement.")
	head1 := flag.String("head1", "", "Head of odd cycle prompt. Fixed statement.")
	head2 := flag.String("head2", "", "Head of even cycle prompt. Fixed statement.")
	tail := flag.String("tail", "", "Head of prompt. Fixed statement.")
	tail1 := flag.String("tail1", "", "Head of odd cycle prompt. Fixed statement.")
	tail2 := flag.String("tail2", "", "Head of even cycle prompt. Fixed statement.")
	flag.Parse()
	prompt := flag.Arg(0)

	var content string
	var err error

	fileName := cmd.GenerateFileName()

	if err := files.Append(fileName, "---\nhead: "+*head+"\nhead1: "+*head1+"\nhead2: "+*head2+
		"\nprompt: "+prompt+
		"\ntail: "+*tail+"\ntail1: "+*tail1+"\ntail2: "+*tail2+"\n---\n"); err != nil {
		log.Fatalf("Error appending to file 1@%d: %v", 0, err)
	}

	client := api.NewClient("http://172.27.167.204:11434")

	for i := 1; i < *cyclesLimit+1; i++ {
		// 整形
		if i%2 != 0 {
			// 1 odd
			prompt = "日本語で出力すること。\n" + *head + "\n" + *head1 + "\n" + prompt + "\n" + *tail + "\n" + *tail1
		} else {
			// 2 even
			prompt = "日本語で出力すること。\n" + *head + "\n" + *head2 + "\n" + prompt + "\n" + *tail + "\n" + *tail2
		}

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

		prompt = content
	}
}
