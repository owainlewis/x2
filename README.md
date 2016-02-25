# X2

An in browser AI that can understand your voice and perform actions based on your commands.

https://x2ai.herokuapp.com/

[![](https://raw.githubusercontent.com/owainlewis/x2/master/static/preview2.png)](https://vimeo.com/154906249 "Speech AI Demo")

A customisable, voice activated intelligent agent for automating tasks. 

X2 can perform complex actions and store information for you.
X2 comes with a voice activated web interface, a simple memory (which can be expanded) and a basic notion of self identity. 

You can deploy X2 on any Linux server. You might want to use it to automate tasks. For example:

> You: Deploy the user service to producton
> X2: Ok. Deploying the user service to production now.

## Deploying your X2 agent to Heroku

```
heroku create
git push heroku master
```

## Creating custom modules

X2 is designed for customisation. A module has two methods Matches and Perform.
The first checks if an input phrase sent to X2 matches
and the second performs some action returning a reply that gets sent back to the caller.

Let's take a look at a very simple module which simply returns the current time when asked
things like "What is the time?" or "What time is it?"

```go
package modules

import (
	"github.com/owainlewis/x2/agent"
	"strings"
	"time"
)

type Time struct {
}

func (t Time) Matches(input string) bool {
	return strings.Contains(input, "what time is it") || strings.Contains(input, "what is the time")
}

func (t Time) Perform(agent *agent.Agent) agent.AgentReply {
	current_time := time.Now().UTC()
	return agent.Reply("The Current time is " + current_time.Format("2006-01-02 MST"))
}
```

After creating a module you need to add add it to your agent

```go
emily := agent.New()
emily.SetActions(modules.Ping{}, modules.Time{}, modules.Weather{})
emily.Query("What is the weather like in San Francisco?")
```
