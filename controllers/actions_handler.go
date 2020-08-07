package controllers

import (
	"encoding/json"
	"runtime/debug"

	socketio "github.com/googollee/go-socket.io"
	"github.com/name5566/leaf/log"
)

// ActionsHandler used to mange callbacks for controllers
type ActionsHandler struct {
	controllerHandler *ControllerHandler
	actions           map[string]func(*Connection) (interface{}, error)
}

// MakeActionsHandler used to instantiate ActionsHandler
func makeActionsHandler(controllerHandler *ControllerHandler) *ActionsHandler {
	actionsHandler := &ActionsHandler{
		controllerHandler: controllerHandler,
		actions:           make(map[string]func(*Connection) (interface{}, error)),
	}

	return actionsHandler
}

// AddAction used to add action
func (handler *ActionsHandler) AddAction(actionName string, action func(*Connection) (interface{}, error)) {
	handler.actions[actionName] = action
	nsp := handler.controllerHandler.nsp

	handler.controllerHandler.server.OnEvent(nsp, actionName, func(conn socketio.Conn, message string) string {
		log.Debug("Receiving Event %v#%v", nsp, actionName)

		var status int
		var errors *Errors
		response, err := handler.call(actionName, conn, message)

		if err != nil {
			status = err.Status

			errors = &Errors{
				Messages: err.Error,
			}
		} else {
			status = 200
		}

		json, _ := json.Marshal(&struct {
			Status   int
			Response interface{} `json:"response"`
			Errors   *Errors
		}{
			Status:   status,
			Response: response,
			Errors:   errors,
		})

		log.Debug("--------------------------------------------")

		return string(json)
	})
}

func (handler *ActionsHandler) call(actionName string, conn socketio.Conn, message string) (response interface{}, networkError *NetworkError) {
	defer func() {
		if r := recover(); r != nil {
			log.Debug("%v", r)
			debug.PrintStack()

			networkError = &NetworkError{
				Status: 500,
				Error:  r.(error).Error(),
			}
		}
	}()

	var params map[string]interface{}
	json.Unmarshal([]byte(message), &params)

	connection := MakeConnection(conn, params)

	err := handler.controllerHandler.callbacksHandler.CallBeforeActions(actionName, connection)

	if err == nil {
		action := handler.actions[actionName]
		response, err = action(connection)
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
