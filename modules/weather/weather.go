package weather

// A module that gets the weather forecast when x2 is asked

import (
	"encoding/json"
	"github.com/owainlewis/x2/agent"
	"net/http"
	"regexp"
)

type weatherValue struct {
	Description string `json:"description"`
}

type weatherData struct {
	Name    string         `json:"name"`
	Weather []weatherValue `json: "weather"`
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=2d00bd3463dc34d772a719b1255bec8d&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

type Weather struct {
}

func (t Weather) Matches(input string) bool {
	var expr = regexp.MustCompile(`the weather`)
	return expr.MatchString(input)
}

func (t Weather) Perform(agent *agent.Agent) agent.AgentReply {

	result, err := query("Cardiff")

	if err != nil {
		return agent.Reply(err.Error())
	}

	description := result.Weather[0].Description

	return agent.Reply("It looks like " + description)
}
