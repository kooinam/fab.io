package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

// Event used as medium of communication
type Event struct {
	createdAt  time.Time
	nsp        string
	room       string
	name       string
	view       interface{}
	parameters map[string]interface{}
}

// makeEvent used to instantiate an socket event
func makeEvent(nsp string, room string, eventName string, view interface{}, parameters map[string]interface{}) *Event {
	formattedNsp := fmt.Sprintf("/%v", nsp)

	return &Event{
		createdAt:  time.Now(),
		nsp:        formattedNsp,
		room:       room,
		name:       eventName,
		view:       view,
		parameters: parameters,
	}
}

// Broadcast used to broadcast event
func (event *Event) Broadcast(server *socketio.Server) {
	json := event.render()

	server.BroadcastToRoom(event.nsp, event.room, event.name, json)
}

func (event *Event) render() string {
	response := make(map[string]interface{})
	response["response"] = event.view
	parameters := event.parameters

	if parameters == nil {
		parameters = make(map[string]interface{})
	}

	response["event"] = &struct {
		CreatedAt  int64                  `json:"createdAt"`
		Name       string                 `json:"name"`
		Parameters map[string]interface{} `json:"parameters"`
	}{
		CreatedAt:  event.createdAt.Unix(),
		Name:       event.name,
		Parameters: parameters,
	}

	json, _ := json.Marshal(response)

	return string(json)
}
