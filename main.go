package main

import (
	"bufio"
	"fmt"
	"github.com/owainlewis/x2/agent"
	"github.com/owainlewis/x2/modules"
	"github.com/owainlewis/x2/persist"
	"github.com/owainlewis/x2/responders"
	"os"
)

func loop(bot *agent.Agent) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> How can I help?: ")
	input, _ := reader.ReadString('\n')
	bot.Tell(input)
	loop(bot)
}

func main() {
	ping := modules.Ping{}
	responder := responders.Console{}
	memory := persist.Transient{}

	bot := agent.New()
	bot.SetName("Emily")
	bot.SetActions(ping)
	bot.SetResponder(responder)
	bot.SetMemory(memory)

	bot.Tell("ping")

	bot.Remember("My name", "owain")

	bot.Recall("My name")
}
