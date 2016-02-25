package main

import (
	"encoding/json"
	"fmt"
	"github.com/owainlewis/x2/agent"
	"github.com/owainlewis/x2/modules"
	"github.com/owainlewis/x2/modules/self"
	"github.com/owainlewis/x2/modules/time"
	"github.com/owainlewis/x2/modules/weather"
	"log"
	"net/http"
	"os"
)

func buildAgent() *agent.Agent {
	emily := agent.New()
	emily.SetName("Emily")

	emily.SetActions(modules.Ping{}, weather.Weather{}, self.Self{}, time.Time{})
	return emily
}

var emily = buildAgent()

func agentRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		query := agent.AgentQuery{}
		err := decoder.Decode(&query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		reply := emily.Query(query)
		fmt.Fprintf(w, reply.Tell)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/agent", agentRequestHandler)
	http.HandleFunc("/", indexHandler)

	log.Println("Starting agent. Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+string(port), nil))
}
