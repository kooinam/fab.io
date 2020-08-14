package actors

// Action is alias for func(context *Context) error
type Action func(context *Context) error

// ActionsHandler used to mange callbacks for controllers
type ActionsHandler struct {
	actions map[string]Action
}

// makeActionsHandler used to instantiate EventsHandler
func makeActionsHandler() *ActionsHandler {
	actionsHandler := &ActionsHandler{
		actions: make(map[string]Action),
	}

	return actionsHandler
}

// RegisterAction used to register action
func (handler *ActionsHandler) RegisterAction(actionName string, action Action) {

}

func (handler *ActionsHandler) handleEvent(event Event) {

}
