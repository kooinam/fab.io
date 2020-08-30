package controllers

import (
	"encoding/json"
	"runtime/debug"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/logger"
)

// Action is alias for func(*Connection) (interface{}, error)
type Action func(*Context)

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

	result := handler.execute(actionName, conn, message)

	json, _ := json.Marshal(&struct {
		Status   string      `json:"string"`
		Response interface{} `json:"response"`
		Error    string      `json:"error"`
	}{
		Status:   result.Status(),
		Response: result.Content(),
		Error:    result.ErrorMessage(),
	})

	logger.Debug("--------------------------------------------")

	return string(json)
}

func (handler *ActionsHandler) execute(actionName string, conn socketio.Conn, message string) (result *Result) {
	defer func() {
		if r := recover(); r != nil {
			logger.Debug("%v", r)
			debug.PrintStack()

			result := makeResult()

			result.Set(nil, StatusInternalServerError, r.(error))
		}
	}()

	var params helpers.H
	json.Unmarshal([]byte(message), &params)

	context := makeContext(conn, params)

	handler.controllerHandler.hooksHandler.executeBeforeActionHooks(actionName, context)

	if context.result != nil {
		action := handler.actions[actionName]

		action(context)

		if context.result == nil {
			context.SetResult(nil, StatusSuccess, nil)
		}
	}

	return context.result
}
