package controllers

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/name5566/leaf/log"
)

var rack *Rack

func init() {
	rack = &Rack{}
}

type Rack struct {
	server             *socketio.Server
	controllerHandlers map[string]*ControllerHandler
}

func Setup() {
	rack.controllerHandlers = make(map[string]*ControllerHandler)

	rack.startServer()
}

// BroadcastEvent used to broadcast event
func BroadcastEvent(nsp string, room string, event string, view interface{}, parameters map[string]interface{}) {
	socketEvent := makeSocketEvent(nsp, room, event, view, parameters)
	json := socketEvent.render()

	rack.server.BroadcastToRoom(socketEvent.nsp, socketEvent.room, socketEvent.name, json)
}

// RegisterController used to register controller
func RegisterController(nsp string, controller Controller) {
	formattedNsp := fmt.Sprintf("/%v", nsp)
	rack.controllerHandlers[formattedNsp] = makeControllerHandler(rack.server, formattedNsp, controller)

	rack.server.OnError(formattedNsp, func(conn socketio.Conn, e error) {
		log.Debug("%v", e)
	})
}

func (rack *Rack) startServer() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Debug("socket.io error %v", err)
	}

	rack.server = server

	server.OnConnect("/", func(conn socketio.Conn) error {
		log.Debug("connected: %v - %v", conn.Namespace(), conn.ID())

		return nil
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		log.Debug("disconnected: %v - %v, %v ", conn.Namespace(), conn.ID(), reason)
	})

	go server.Serve()

	go func() {
		http.Handle("/socket.io/", server)

		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))

		log.Debug("starting Socket Server...")

		http.ListenAndServe(":8000", nil)
	}()
}
