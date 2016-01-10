package agent

type Agent struct {
	Name    string
	Actions []Action
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

func (agent *Agent) Act(input string) {
	for _, action := range agent.Actions {
		if action.Matches(input) {
			action.Perform()
			return
		}
	}
}

func New() *Agent {
	return &Agent{}
}
