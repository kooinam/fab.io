package fab

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fabio/controllers"
	"github.com/kooinam/fabio/logger"
)

var engine *Engine

// Engine is the core rack for fab.io
type Engine struct {
	server             *socketio.Server
	controllerHandlers map[string]*controllers.ControllerHandler
}

// Setup used to setup engine
func Setup() {
	engine = &Engine{
		controllerHandlers: make(map[string]*controllers.ControllerHandler),
	}

	engine.setup()
}

// Serve used to serve
func Serve(httpHandler func()) {
	server := engine.server

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
func BroadcastEvent(nsp string, room string, event string, view interface{}, parameters map[string]interface{}) {
	socketEvent := controllers.MakeSocketEvent(nsp, room, event, view, parameters)

	socketEvent.Broadcast(engine.server)
}

// RegisterController used to register controller
func RegisterController(nsp string, controller controllers.Controller) {
	formattedNsp := fmt.Sprintf("/%v", nsp)
	engine.controllerHandlers[formattedNsp] = controllers.MakeControllerHandler(engine.server, formattedNsp, controller)

	engine.server.OnError(formattedNsp, func(conn socketio.Conn, e error) {
		logger.Debug("%v", e)
	})
}

func (engine *Engine) setup() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		logger.Debug("socket.io error %v", err)
	}

	engine.server = server

	server.OnConnect("/", func(conn socketio.Conn) error {
		logger.Debug("connected: %v - %v", conn.Namespace(), conn.ID())

		return nil
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		logger.Debug("disconnected: %v - %v, %v ", conn.Namespace(), conn.ID(), reason)
	})
}
