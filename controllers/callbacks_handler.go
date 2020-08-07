package controllers

// CallbacksHandler used to mange callbacks for controllers
type CallbacksHandler struct {
	beforeActions []func(string, *Connection) error
}

// MakeCallbacksHandler used to instantiate CallbacksHandler
func makeCallbacksHandler() *CallbacksHandler {
	callbacksHandler := &CallbacksHandler{
		beforeActions: []func(string, *Connection) error{},
	}

	return callbacksHandler
}

// AddBeforeAction used to add before action callback
func (handler *CallbacksHandler) AddBeforeAction(beforeAction func(string, *Connection) error) {
	handler.beforeActions = append(handler.beforeActions, beforeAction)
}

// CallBeforeActions used to call before actions callback
func (handler *CallbacksHandler) CallBeforeActions(action string, connection *Connection) error {
	var err error

	for _, callback := range handler.beforeActions {
		err = callback(action, connection)

		if err != nil {
			break
		}
	}

	return err
}
