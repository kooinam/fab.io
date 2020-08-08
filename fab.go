package fab

import (
	"fabio/controllers"
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/name5566/leaf/log"
)

var engine *Engine

// Engine is the core rack for fab.io
type Engine struct {
	server             *socketio.Server
	controllerHandlers map[string]*controllers.ControllerHandler
}

// Setup used to update setup
func Setup() {
	fmt.Printf("hello...ddd")

	engine.controllerHandlers = make(map[string]*controllers.ControllerHandler)

	engine.startServer()
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
		log.Debug("%v", e)
	})
}

func (engine *Engine) startServer() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Debug("socket.io error %v", err)
	}

	engine.server = server

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
