package controllers

import (
	"encoding/json"
	"runtime/debug"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/logger"
)

// Hook is alias for func(string, *Context)
type Hook func(string, *Context)

// Action is alias for func(*Connection) (interface{}, error)
type Action func(*Context)

// ActionsHandler used to mange callbacks for controllers
type ActionsHandler struct {
	manager           *Manager
	controllerHandler *ControllerHandler
	beforeActionHooks []Hook
	actions           map[string]Action
}

// makeActionsHandler used to instantiate ActionsHandler
func makeActionsHandler(manager *Manager, controllerHandler *ControllerHandler) *ActionsHandler {
	actionsHandler := &ActionsHandler{
		manager:           manager,
		controllerHandler: controllerHandler,
		beforeActionHooks: []Hook{},
		actions:           make(map[string]Action),
	}

	return actionsHandler
}

// RegisterBeforeActionHook used to add before hook
func (handler *ActionsHandler) RegisterBeforeActionHook(beforeActionHook Hook) {
	handler.beforeActionHooks = append(handler.beforeActionHooks, beforeActionHook)
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
		context := makeContext(handler.manager.viewsManager, conn, helpers.H{})

		action(context)

		return nil
	})
}

// RegisterDisconnectedAction used to register disconnected action
func (handler *ActionsHandler) RegisterDisconnectedAction(action Action) {
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnDisconnect(nsp, func(conn socketio.Conn, reason string) {
		context := makeContext(handler.manager.viewsManager, conn, helpers.H{
			"reason": reason,
		})

		action(context)
	})
}

// RegisterErrorAction used to register error action
func (handler *ActionsHandler) RegisterErrorAction(action Action) {
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnError(nsp, func(conn socketio.Conn, err error) {
		context := makeContext(handler.manager.viewsManager, conn, helpers.H{
			"error": err.Error(),
		})

		action(context)
	})
}

// executeBeforeActionHooks used to execute before action hooks
func (handler *ActionsHandler) executeBeforeActionHooks(action string, context *Context) {
	for _, hook := range handler.beforeActionHooks {
		hook(action, context)

		if context.result != nil {
			break
		}
	}
}

func (handler *ActionsHandler) handleAction(nsp string, actionName string, conn socketio.Conn, message string) string {
	logger.Debug("Receiving Event: %v#%v Message: %v", nsp, actionName, message)

	result := handler.execute(actionName, conn, message)

	json, _ := json.Marshal(&struct {
		Status   string      `json:"status"`
		Response interface{} `json:"response,omitempty"`
		Error    string      `json:"error,omitempty"`
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

	context := makeContext(handler.manager.viewsManager, conn, params)

	handler.executeBeforeActionHooks(actionName, context)

	if context.result == nil {
		action := handler.actions[actionName]

		action(context)

		if context.result == nil {
			context.SetSuccessResult(nil)
		}
	}

	return context.result
}
