package self

// This package contains some basic stuff that underpins all agents
// It includes a notion of self

import (
	"github.com/owainlewis/x2/agent"
	"regexp"
)

type Self struct {
}

func greeting(input string) bool {
	var expr = regexp.MustCompile(`hello|hi|hey`)
	return expr.MatchString(input)
}
func feeling(input string) bool {
	var expr = regexp.MustCompile(`how are you feeling`)
	return expr.MatchString(input)
}

func (self Self) Matches(input string) bool {
	if greeting(input) {
		return true
	}
	if feeling(input) {
		return true
	}
	return false
}

func (self Self) Perform(agent *agent.Agent) agent.AgentReply {
	if greeting(agent.LastQuery.Ask) {
		return agent.Reply("Hi. My name is " + agent.Name)
	}
	if feeling(agent.LastQuery.Ask) {
		return agent.Reply("I'm great. Thanks for asking")
	}
	return agent.Reply("Pardon?")
}
