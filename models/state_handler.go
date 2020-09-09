package models

// StateHandler used to handle fsm's callback
type StateHandler struct {
	name      string
	enterHook func(string)
	runHook   func(float64)
	exitHook  func()
}

func makeStateHandler(name string) *StateHandler {
	handler := &StateHandler{
		name: name,
	}

	return handler
}

// WithEnterHook used to register state handler's enter hook
func (handler *StateHandler) WithEnterHook(hook func(string)) *StateHandler {
	handler.enterHook = hook

	return handler
}

// WithRunHook used to register state handler's run hook
func (handler *StateHandler) WithRunHook(hook func(float64)) *StateHandler {
	handler.runHook = hook

	return handler
}

// WithExitHook used to register state handler's exit hook
func (handler *StateHandler) WithExitHook(hook func()) *StateHandler {
	handler.exitHook = hook

	return handler
}
