package main

import (
	"github.com/owainlewis/x2/agent"
	"github.com/owainlewis/x2/modules"
	"github.com/owainlewis/x2/responders"
)

func main() {

	ping := modules.Ping{}
	responder := responders.Console{}

	bot := agent.New()
	bot.SetName("owain")
	bot.AddAction(ping)
	bot.SetResponder(responder)

	bot.Tell("Hello. What is your name?")
	bot.Tell("ping")
}
