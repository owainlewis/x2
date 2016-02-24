package self

// This package contains some basic stuff that underpins all agents
// It includes a notion of self

import (
	"github.com/owainlewis/x2/agent"
)

type Self struct {
}

func (chat Self) Matches(input string) bool {
	if input == "who are you" {
		return true
	}
	return false
}

func (chat Self) Perform(agent *agent.Agent) agent.AgentReply {
	println(agent.LastQuery.Ask)
	return agent.Reply("My name is " + agent.Name)
}
