package x2

type Bot struct {
	Name string
}

func (bot *Bot) SetName(name string) {
	bot.Name = name
}
