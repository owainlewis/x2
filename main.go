package main

import (
	"fmt"
	"github.com/owainlewis/x2/agent"
	"github.com/owainlewis/x2/modules"
)

func main() {

	ping := modules.Ping{}

	bot := agent.New()
	bot.SetName("owain")
	bot.AddAction(ping)
	fmt.Println(bot.Name)

	bot.Act("ping")

}
