package main

import (
	"fmt"
	"github.com/owainlewis/x2"
)

func main() {
	bot := x2.Bot{}
	bot.SetName("owain")
	fmt.Println(bot.Name)
}
