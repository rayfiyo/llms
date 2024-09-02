package files

import (
	"fmt"

	"github.com/rayfiyo/llms/dialogue/cmd/flags"
)

func Header(fileName string, prompt string) error {
	return Append(fileName, "---"+
		"\nn: "+fmt.Sprint(*flags.CyclesLimit)+
		"\n"+
		"\nModel: "+*flags.Model+
		"\nModel1: "+*flags.Model1+
		"\nModel2: "+*flags.Model2+
		"\n"+
		"\nHead: "+*flags.Head+
		"\nHead1: "+*flags.Head1+
		"\nHead2: "+*flags.Head2+
		"\n"+
		"\nprompt: "+prompt+
		"\nInit: "+*flags.Init+
		"\n"+
		"\nTail: "+*flags.Tail+
		"\nTail1: "+*flags.Tail1+
		"\nTail2: "+*flags.Tail2+
		"\n---\n",
	)
}
