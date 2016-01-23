package modules

import (
	"github.com/owainlewis/x2/agent"
	"strings"
)

type Ping struct {
}

func (ping Ping) Matches(input string) bool {
	return strings.Contains(input, "ping") || strings.Contains(input, "PING")
}

func (ping Ping) Perform(agent *agent.Agent) agent.AgentReply {
	return agent.Reply("PONG")
}
