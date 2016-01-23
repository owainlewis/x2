package agent

type AgentQuery struct {
	Ask string
}

type AgentReply struct {
	Tell string
}

// An action is performed by an agent.
// When the agent performs an action it must return a reply back to the caller
type Action interface {
	Matches(input string) bool
	Perform(agent *Agent) AgentReply
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
	// Linear search over actions
	for _, action := range agent.Actions {
		if action.Matches(input) {
			return true
		}
	}
	return false
}

func (agent *Agent) Reply(what string) AgentReply {
	return AgentReply{what}
}

func (agent *Agent) Query(query AgentQuery) AgentReply {
	// Linear search over actions
	if query.Ask == "" {
		return agent.Reply("Your query was empty")
	}
	for _, action := range agent.Actions {
		if action.Matches(query.Ask) {
			// an action needs a reference to the underlying agent
			// this is needed to query agent memory or identity
			return action.Perform(agent)
		}
	}
	return agent.Reply("Sorry. I don't understand")
}

func New() *Agent {
	return &Agent{}
}
