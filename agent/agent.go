package agent

type AgentQuery struct {
	Ask string
}

type AgentReply struct {
	Tell string
}

type Agent struct {
	Name    string
	Actions []Action
}

func (agent *Agent) SetName(name string) {
	agent.Name = name
}

func (agent *Agent) SetActions(actions ...Action) {
	for _, action := range actions {
		agent.Actions = append(agent.Actions, action)
	}
}

func (agent *Agent) Understands(input string) bool {
	for _, action := range agent.Actions {
		if action.Matches(input) {
			return true
		}
	}
	return false
}

func (agent *Agent) Query(command AgentQuery) AgentReply {
	return AgentReply{"Affirmative"}
}

func New() *Agent {
	return &Agent{}
}
