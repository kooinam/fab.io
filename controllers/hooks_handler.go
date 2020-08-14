package controllers

// Hook is alias for Hook
type Hook func(string, *Context) error

// HooksHandler used to mange callbacks for controllers
type HooksHandler struct {
	beforeHooks []Hook
}

// makeHooksHandler used to instantiate CallbacksHandler
func makeHooksHandler() *HooksHandler {
	hooksHandler := &HooksHandler{
		beforeHooks: []Hook{},
	}

	return hooksHandler
}

// RegisterBeforeHook used to add before hook
func (handler *HooksHandler) RegisterBeforeHook(beforeHook Hook) {
	handler.beforeHooks = append(handler.beforeHooks, beforeHook)
}

// executeBeforeHooks used to execute before hooks
func (handler *HooksHandler) executeBeforeHooks(action string, context *Context) error {
	var err error

	for _, hook := range handler.beforeHooks {
		err = hook(action, context)

		if err != nil {
			break
		}
	}

	return err
}
