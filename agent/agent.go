package agent

// An agent is something that identity, memory and can perform actions to given user inputs
type Agent struct {
	Name    string
	Actions []Action
	Respond Responder
}

func (agent *Agent) SetName(name string) {
	agent.Name = name
}

func (agent *Agent) AddAction(action Action) {
	agent.Actions = append(agent.Actions, action)
}

func (agent *Agent) SetActions(actions []Action) {
	agent.Actions = actions
}

func (agent *Agent) SetResponder(responder Responder) {
	agent.Respond = responder
}

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

func New() *Agent {
	return &Agent{}
}
