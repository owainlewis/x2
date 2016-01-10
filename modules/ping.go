package modules

import "github.com/owainlewis/x2/agent"

type Ping struct {
}

func (ping Ping) Matches(input string) bool {
	return input == "ping"
}

func (ping Ping) Perform(agent *agent.Agent) {
	agent.Say("PONG")
}
