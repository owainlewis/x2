package agent

type Action interface {
	Matches(input string) bool
	Perform(agent *Agent)
}
