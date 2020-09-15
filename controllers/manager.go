package controllers

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/logger"
)

// Manager is singleton manager for controller module
type Manager struct {
	server             *socketio.Server
	controllerHandlers map[string]*ControllerHandler
}

// Setup used to setup cotroller manager
func (manager *Manager) Setup() {
	manager.controllerHandlers = make(map[string]*ControllerHandler)

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
func (manager *Manager) Serve(port string, httpHandler func()) {
	logger.Debug("Initializing fab.io...")

	server := manager.server

	http.Handle("/socket.io/", corsMiddleware(server))

	if httpHandler != nil {
		httpHandler()
	}

	go server.Serve()

	logger.Debug("Starting Socket Server @ %v...", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

// BroadcastEvent used to broadcast event
func (manager *Manager) BroadcastEvent(nsp string, room string, eventName string, view interface{}, parameters helpers.H) {
	event := makeEvent(nsp, room, eventName, view, parameters)

	event.Broadcast(manager.server)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}
