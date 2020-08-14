package controllers

import (
	socketio "github.com/googollee/go-socket.io"
)

// ControllerHandler used to handle controller
type ControllerHandler struct {
	server         *socketio.Server
	nsp            string
	controller     Controllable
	hooksHandler   *HooksHandler
	actionsHandler *ActionsHandler
}

// makeControllerHandler used to instantiate controller handler
func makeControllerHandler(server *socketio.Server, nsp string, controllable Controllable) *ControllerHandler {
	handler := &ControllerHandler{
		server:       server,
		nsp:          nsp,
		controller:   controllable,
		hooksHandler: makeHooksHandler(),
	}
	handler.actionsHandler = makeActionsHandler(handler)

	handler.controller.RegisterBeforeHooks(handler.hooksHandler)
	handler.controller.RegisterActions(handler.actionsHandler)

	return handler
}
