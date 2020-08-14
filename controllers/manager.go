package controllers

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fabio/logger"
)

// Manager is singleton manager for controller module
type Manager struct {
	server             *socketio.Server
	controllerHandlers map[string]*ControllerHandler
}

// Setup used to setup cotroller manager
func (manager *Manager) Setup() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		logger.Debug("socket.io error %v", err)
	}

	manager.server = server

	server.OnConnect("/", func(conn socketio.Conn) error {
		logger.Debug("connected: %v - %v", conn.Namespace(), conn.ID())

		return nil
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		logger.Debug("disconnected: %v - %v, %v ", conn.Namespace(), conn.ID(), reason)
	})
}

// RegisterController used to register controller
func (manager *Manager) RegisterController(nsp string, controllable Controllable) {
	formattedNsp := fmt.Sprintf("/%v", nsp)
	manager.controllerHandlers[formattedNsp] = makeControllerHandler(manager.server, formattedNsp, controllable)

	manager.server.OnError(formattedNsp, func(conn socketio.Conn, e error) {
		logger.Debug("%v", e)
	})
}

// Serve used to serve
func (manager *Manager) Serve(httpHandler func()) {
	server := manager.server

	go server.Serve()

	logger.Debug("Initializing fab.io...")

	http.Handle("/socket.io/", server)

	if httpHandler != nil {
		httpHandler()
	}

	logger.Debug("Starting Socket Server...")

	http.ListenAndServe(":8000", nil)
}

// BroadcastEvent used to broadcast event
func (manager *Manager) BroadcastEvent(nsp string, room string, eventName string, view interface{}, parameters map[string]interface{}) {
	event := makeEvent(nsp, room, eventName, view, parameters)

	event.Broadcast(manager.server)
}
