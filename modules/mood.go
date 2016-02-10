package modules

import (
	"github.com/owainlewis/x2/agent"
)

type Mood struct {
}

func (ping Mood) Matches(input string) bool {
	if input == "how are you feeling" {
		return true
	}
	return false
}

func (ping Mood) Perform(agent *agent.Agent) agent.AgentReply {
	return agent.Reply("I am feeling great. Thanks for asking")
}
