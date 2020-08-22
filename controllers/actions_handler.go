package controllers

import (
	"encoding/json"
	"runtime/debug"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/logger"
)

// Action is handler for func(*Connection) (interface{}, error)
type Action func(*Context) (interface{}, error)

// ActionsHandler used to mange callbacks for controllers
type ActionsHandler struct {
	controllerHandler *ControllerHandler
	actions           map[string]Action
}

// makeActionsHandler used to instantiate ActionsHandler
func makeActionsHandler(controllerHandler *ControllerHandler) *ActionsHandler {
	actionsHandler := &ActionsHandler{
		controllerHandler: controllerHandler,
		actions:           make(map[string]Action),
	}

	return actionsHandler
}

// RegisterAction used to register action
func (handler *ActionsHandler) RegisterAction(actionName string, action Action) {
	handler.actions[actionName] = action
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnEvent(nsp, actionName, func(conn socketio.Conn, message string) string {
		return handler.handleAction(nsp, actionName, conn, message)
	})
}

// RegisterConnectedAction used to register connected action
func (handler *ActionsHandler) RegisterConnectedAction(action Action) {
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnConnect(nsp, func(conn socketio.Conn) error {
		context := makeContext(conn, helpers.H{})

		action(context)

		return nil
	})
}

// RegisterDisconnectedAction used to register disconnected action
func (handler *ActionsHandler) RegisterDisconnectedAction(action Action) {
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnDisconnect(nsp, func(conn socketio.Conn, reason string) {
		context := makeContext(conn, helpers.H{
			"reason": reason,
		})

		action(context)
	})
}

// RegisterErrorAction used to register error action
func (handler *ActionsHandler) RegisterErrorAction(action Action) {
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnError(nsp, func(conn socketio.Conn, err error) {
		context := makeContext(conn, helpers.H{
			"error": err.Error(),
		})

		action(context)
	})
}

func (handler *ActionsHandler) handleAction(nsp string, actionName string, conn socketio.Conn, message string) string {
	logger.Debug("Receiving Event %v#%v", nsp, actionName)

	var status int
	var errorsView *ErrorsView
	response, err := handler.execute(actionName, conn, message)

	if err != nil {
		status = err.Status

		errorsView = &ErrorsView{
			Messages: err.Error,
		}
	} else {
		status = 200
	}

	json, _ := json.Marshal(&struct {
		Status   int
		Response interface{} `json:"response"`
		Errors   *ErrorsView
	}{
		Status:   status,
		Response: response,
		Errors:   errorsView,
	})

	logger.Debug("--------------------------------------------")

	return string(json)
}

func (handler *ActionsHandler) execute(actionName string, conn socketio.Conn, message string) (response interface{}, networkError *NetworkError) {
	defer func() {
		if r := recover(); r != nil {
			logger.Debug("%v", r)
			debug.PrintStack()

			networkError = &NetworkError{
				Status: 500,
				Error:  r.(error).Error(),
			}
		}
	}()

	var params helpers.H
	json.Unmarshal([]byte(message), &params)

	context := makeContext(conn, params)

	err := handler.controllerHandler.hooksHandler.executeBeforeHooks(actionName, context)

	if err == nil {
		action := handler.actions[actionName]
		response, err = action(context)
	}

	if err != nil {
		// Validation failed
		networkError = &NetworkError{
			Status: 422,
			Error:  err.Error(),
		}
	}

	return response, networkError
}
