package controllers

import (
	"encoding/json"
	"fmt"
	"time"
)

// SocketEvent used as medium of communication
type SocketEvent struct {
	createdAt  time.Time
	nsp        string
	room       string
	name       string
	view       interface{}
	parameters map[string]interface{}
}

func makeSocketEvent(nsp string, room string, eventName string, view interface{}, parameters map[string]interface{}) *SocketEvent {
	formattedNsp := fmt.Sprintf("/%v", nsp)

	return &SocketEvent{
		createdAt:  time.Now(),
		nsp:        formattedNsp,
		room:       room,
		name:       eventName,
		view:       view,
		parameters: parameters,
	}
}

func (socketEvent *SocketEvent) render() string {
	response := make(map[string]interface{})
	response["response"] = socketEvent.view
	parameters := socketEvent.parameters

	if parameters == nil {
		parameters = make(map[string]interface{})
	}

	response["event"] = &struct {
		CreatedAt  int64                  `json:"createdAt"`
		Name       string                 `json:"name"`
		Parameters map[string]interface{} `json:"parameters"`
	}{
		CreatedAt:  socketEvent.createdAt.Unix(),
		Name:       socketEvent.name,
		Parameters: parameters,
	}

	json, _ := json.Marshal(response)

	return string(json)
}
