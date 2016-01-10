package modules

type Ping struct {
}

func (ping Ping) Matches(input string) bool {
	return input == "ping"
}

func (ping Ping) Perform() {
	println("PONG")
}
