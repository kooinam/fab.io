package controllers

// Hook is alias for Hook
type Hook func(string, *Context)

// HooksHandler used to mange callbacks for controllers
type HooksHandler struct {
	beforeActionHooks []Hook
}

// makeHooksHandler used to instantiate CallbacksHandler
func makeHooksHandler() *HooksHandler {
	hooksHandler := &HooksHandler{
		beforeActionHooks: []Hook{},
	}

	return hooksHandler
}

// RegisterBeforeActionHook used to add before hook
func (handler *HooksHandler) RegisterBeforeActionHook(beforeActionHook Hook) {
	handler.beforeActionHooks = append(handler.beforeActionHooks, beforeActionHook)
}

// executeBeforeActionHooks used to execute before action hooks
func (handler *HooksHandler) executeBeforeActionHooks(action string, context *Context) {
	for _, hook := range handler.beforeActionHooks {
		hook(action, context)

		if context.result != nil {
			break
		}
	}
}
