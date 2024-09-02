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

	if err := files.Append(fileName, "---"+
		"\nflags.Head: "+*flags.Head+
		"\nflags.Head1: "+*flags.Head1+
		"\nflags.Head2: "+*flags.Head2+
		"\nprompt: "+prompt+
		"\nflags.Init: "+*flags.Init+
		"\nflags.Tail: "+*flags.Tail+
		"\nflags.Tail1: "+*flags.Tail1+
		"\nflags.Tail2: "+*flags.Tail2+
		"\n---\n",
	); err != nil {
		log.Fatalf("Error appending options to file: %v", err)
	}

	client := api.NewClient("http://172.27.167.204:11434")

	var content string
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

		switch *flags.Mode {
		case "chat":
			request := &models.ChatRequest{
				Model: *flags.Model,
				Messages: []models.Message{
					{Role: "user", Content: prompt},
				},
			}
			content, err = client.Chat(request)
		case "generate":
			request := &models.GenerateRequest{
				Model:  *flags.Model,
				Prompt: prompt,
			}
			content, err = client.Generate(request)
		default:
			log.Fatalf(
				"Invalid flags.Mode: %s. Use 'chat' or 'generate'", *flags.Mode,
			)
		}

		if err != nil {
			log.Fatalf("Error in switch@%d: %v", i, err)
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
