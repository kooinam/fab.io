package controllers

import (
	socketio "github.com/googollee/go-socket.io"
)

type Errors struct {
	Messages string
}

type NetworkError struct {
	Status int
	Error  string
}

// ControllerHandler used to handle controller
type ControllerHandler struct {
	server           *socketio.Server
	nsp              string
	controller       Controller
	callbacksHandler *CallbacksHandler
	actionsHandler   *ActionsHandler
}

// MakeControllerHandler used to instantiate controller handler
func MakeControllerHandler(server *socketio.Server, nsp string, controller Controller) *ControllerHandler {
	handler := &ControllerHandler{
		server:           server,
		nsp:              nsp,
		controller:       controller,
		callbacksHandler: makeCallbacksHandler(),
	}
	handler.actionsHandler = makeActionsHandler(handler)

	handler.controller.AddBeforeActions(handler.callbacksHandler)
	handler.controller.AddActions(handler.actionsHandler)

	return handler
}
