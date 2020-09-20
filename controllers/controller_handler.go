package controllers

import (
	socketio "github.com/googollee/go-socket.io"
)

// ControllerHandler used to handle controller
type ControllerHandler struct {
	manager        *Manager
	server         *socketio.Server
	nsp            string
	controller     Controllable
	actionsHandler *ActionsHandler
}

// makeControllerHandler used to instantiate controller handler
func makeControllerHandler(manager *Manager, server *socketio.Server, nsp string, controllable Controllable) *ControllerHandler {
	handler := &ControllerHandler{
		manager:    manager,
		server:     server,
		nsp:        nsp,
		controller: controllable,
	}
	handler.actionsHandler = makeActionsHandler(manager, handler)

	handler.controller.RegisterActions(handler.actionsHandler)

	return handler
}
