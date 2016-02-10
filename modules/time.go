package modules

import (
	"github.com/owainlewis/x2/agent"
	"strings"
	"time"
)

type Time struct {
}

func (t Time) Matches(input string) bool {
	return strings.Contains(input, "what time is it") ||
		strings.Contains(input, "what is the time")
}

func (t Time) Perform(agent *agent.Agent) agent.AgentReply {
	current_time := time.Now().UTC()
	return agent.Reply("The Current time is " + current_time.Format("2006-01-02 MST"))
}
