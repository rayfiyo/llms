package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rayfiyo/llms/dialogue/cmd/api"
	"github.com/rayfiyo/llms/dialogue/cmd/files"
	"github.com/rayfiyo/llms/dialogue/cmd/flags"
	"github.com/rayfiyo/llms/dialogue/cmd/generate"
	"github.com/rayfiyo/llms/dialogue/models"
)

const language = ""

// const language = "日本語のみで出力すること。"

func main() {
	flags.Parse()
	prompt := flag.Arg(0)

	fileName := generate.FileName()

	if err := files.Header(fileName, prompt); err != nil {
		log.Fatal("Error writing options in file header: %w", err)
	}

	client := api.NewClient("http://172.27.167.204:11434")

	var content string
	var context []int
	var err error

	for i := 1; i < *flags.CyclesLimit+1; i++ {
		// 整形
		if i%2 != 0 {
			// 1 odd
			prompt = language + "\n" +
				*flags.Head + "\n" +
				*flags.Head1 + "\n" +
				prompt + "\n" +
				*flags.Tail + "\n" +
				*flags.Tail1 + "\n"
			if *flags.Model1 != "" {
				flags.Model = flags.Model1
			}
		} else {
			// 2 even
			prompt = language + "\n" +
				*flags.Head + "\n" +
				*flags.Head2 + "\n" +
				prompt + "\n" +
				*flags.Tail + "\n" +
				*flags.Tail2 + "\n"
			if *flags.Model2 != "" {
				flags.Model = flags.Model2
			}
		}
		if *flags.Init != "" {
			prompt = language + "\n" +
				*flags.Head + "\n" +
				*flags.Init + "\n" +
				*flags.Tail + "\n"
			i = 0
			log.Println(*flags.Init)
			*flags.Init = ""
			log.Println(*flags.Init)
		}

		fmt.Print("\n- - - - - - - - - - - -\n")
		log.Printf("%3d:\n\n", i)

		switch *flags.Mode {
		case "chat":
			request := &models.ChatRequest{
				Model: *flags.Model,
				Messages: []models.Message{
					{Role: "user", Content: prompt},
				},
			}
			content, _, err = client.Chat(request)
		case "generate":
			request := &models.GenerateRequest{
				Model:   *flags.Model,
				Prompt:  prompt,
				Context: context,
			}
			content, context, err = client.Generate(request)
		default:
			log.Fatalf(
				"Invalid flags.Mode: %s. Use 'chat' or 'generate'", *flags.Mode,
			)
		}

		if err != nil {
			log.Fatalf("Error in switch@%d: %v", i, err)
		}

		// ファイルに保存
		if err := files.Append(fileName, "## "+fmt.Sprint(i)); err != nil {
			log.Fatalf("Error appending to file 1@%d: %v", i, err)
		}
		if err := files.Append(fileName, content); err != nil {
			log.Fatalf("Error appending to file 2@%d: %v", i, err)
		}

		// 出力を入力に繋げる
		prompt = content
	}
}
