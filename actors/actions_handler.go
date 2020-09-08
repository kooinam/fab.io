package actors

// Action is alias for func(context *Context) error
type Action func(context *Context) error

// ActionsHandler used to mange callbacks for controllers
type ActionsHandler struct {
	identifier string
	actions    map[string]Action
}

// makeActionsHandler used to instantiate EventsHandler
func makeActionsHandler(identifier string) *ActionsHandler {
	actionsHandler := &ActionsHandler{
		identifier: identifier,
		actions:    make(map[string]Action),
	}

	return actionsHandler
}

// RegisterAction used to register action
func (handler *ActionsHandler) RegisterAction(actionName string, action Action) {
	if handler.actions[actionName] != nil {
		panic("actor action registered")
	}

	handler.actions[actionName] = action
}

func (handler *ActionsHandler) handleEvent(event event) {
	action := handler.actions[event.name]

	if action != nil {
		context := makeContext(event.params)
		err := action(context)

		if err != nil {
			event.nak(err.Error())
		} else {
			event.ack()
		}
	} else {
		event.nak("no action found")
	}
}
