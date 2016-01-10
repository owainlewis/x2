package responders

import (
	"bufio"
	"fmt"
	"os"
)

type Console struct {
}

func (responder Console) Ask(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	return text
}

func (responder Console) Tell(input string) {
	fmt.Println(input)
}
