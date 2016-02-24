package weather

import (
	"github.com/owainlewis/x2/agent"
	"strings"
)

type Weather struct {
}

func (t Weather) Matches(input string) bool {
	return strings.Contains(input, "how's the weather")
}

func (t Weather) Perform(agent *agent.Agent) agent.AgentReply {
	return agent.Reply("Rainy with a change of more rain")
}
