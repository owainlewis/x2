package agent

// An agent is something that identity, memory and can perform actions to given user inputs
type Agent struct {
	Name    string
	Actions []Action
	Respond Responder
	Memory  Memory
}

// Init functions
// ************************************************************************************

func (agent *Agent) SetName(name string) {
	agent.Name = name
}

func (agent *Agent) SetActions(actions ...Action) {
	for _, action := range actions {
		agent.Actions = append(agent.Actions, action)
	}
}

func (agent *Agent) SetResponder(responder Responder) {
	agent.Respond = responder
}

func (agent *Agent) SetMemory(memory Memory) {
	agent.Memory = memory
}

// Actions
// ************************************************************************************

func (agent *Agent) Understands(input string) bool {
	for _, action := range agent.Actions {
		if action.Matches(input) {
			return true
		}
	}
	return false
}

func (agent *Agent) Tell(input string) {
	if agent.Understands(input) {
		for _, action := range agent.Actions {
			if action.Matches(input) {
				action.Perform(agent)
				return
			}
		}
	} else {
		agent.Say("Sorry. I don't understand")
	}
}

func (agent *Agent) Say(output string) {
	agent.Respond.Tell(output)
}

func (agent *Agent) Remember(k string, v string) {
	agent.Memory.Set(k, v)
}

func (agent *Agent) Recall(k string) {
	recalled := agent.Memory.Get(k)
	if recalled != "" {
		agent.Say(recalled)
	} else {
		agent.Say("I don't recall")
	}
}

// Construct
// ************************************************************************************

func New() *Agent {
	return &Agent{}
}
