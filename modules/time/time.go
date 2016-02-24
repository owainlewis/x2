package time

import (
	"github.com/owainlewis/x2/agent"
	"regexp"
	"time"
)

type Time struct {
}

func (t Time) Matches(input string) bool {
	var expr = regexp.MustCompile(`time`)
	return expr.MatchString(input)
}

func (t Time) Perform(agent *agent.Agent) agent.AgentReply {
	current_time := time.Now().UTC()
	return agent.Reply("The Current time is " + current_time.Format("2006-01-02 MST"))
}
