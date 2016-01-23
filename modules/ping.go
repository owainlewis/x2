package modules

import (
	"github.com/owainlewis/x2/agent"
	"log"
	"strings"
)

type Ping struct {
}

func (ping Ping) Matches(input string) bool {
	log.Printf(input)
	return strings.Contains(input, "ping")
}

func (ping Ping) Perform(agent *agent.Agent) agent.AgentReply {
	return agent.Reply("PONG")
}
