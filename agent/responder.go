package agent

type Responder interface {
	Tell(input string)
}
