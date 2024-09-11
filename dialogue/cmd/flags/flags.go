package flags

import (
	"flag"
)

var (
	Mode = flag.String(
		"mode", "chat", "Mode to use: 'chat' or 'generate'.")
	Model = flag.String(
		"model", "Llama-3-Swallow-70B-Instruct-v0.1-Q8_0", "Set model name.")
	Model1 = flag.String(
		"model1", "", "Overwrite model name for odd cycles.")
	Model2 = flag.String(
		"model2", "", "Overwrite model name for even cycles.")
	CyclesLimit = flag.Int(
		"n", 6, "Limit number of sends cycles.")
	Head = flag.String(
		"head", "", "Head of prompt. Fixed statement.")
	Head1 = flag.String(
		"head1", "", "Head of odd cycle prompt. Fixed statement.")
	Head2 = flag.String(
		"head2", "", "Head of even cycle prompt. Fixed statement.")
	Tail = flag.String(
		"tail", "", "Head of prompt. Fixed statement.")
	Tail1 = flag.String(
		"tail1", "", "Head of odd cycle prompt. Fixed statement.")
	Tail2 = flag.String(
		"tail2", "", "Head of even cycle prompt. Fixed statement.")
	Init = flag.String(
		"init", "", "Set prompt as initial question (0th).")
)

func Parse() {
	flag.Parse()
}
