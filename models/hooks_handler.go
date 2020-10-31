package models

import (
	"github.com/kooinam/fab.io/helpers"
)

// Validator is alias for func(*Connection) (interface{}, error)
type Validator func() error

// HooksHandler used to mange callbacks for models
type HooksHandler struct {
	initializeHook    func(*helpers.Dictionary)
	validationHooks   []Validator
	afterCreateHooks  []func()
	afterDestroyHooks []func()
}

// makeHooksHandler used to instantiate CallbacksHandler
func makeHooksHandler() *HooksHandler {
	hooksHandler := &HooksHandler{
		validationHooks:   []Validator{},
		afterCreateHooks:  []func(){},
		afterDestroyHooks: []func(){},
	}

	return hooksHandler
}

// RegisterInitializeHook used to add initialize hook
func (handler *HooksHandler) RegisterInitializeHook(initializeHook func(*helpers.Dictionary)) {
	handler.initializeHook = initializeHook
}

// RegisterValidationHook used to add a validation hook
func (handler *HooksHandler) RegisterValidationHook(validationHook Validator) {
	handler.validationHooks = append(handler.validationHooks, validationHook)
}

// RegisterAfterCreateHook used to add after create hook
func (handler *HooksHandler) RegisterAfterCreateHook(afterCreateHook func()) {
	handler.afterCreateHooks = append(handler.afterCreateHooks, afterCreateHook)
}

// RegisterAfterDestroyHook used to add after destroy hook
func (handler *HooksHandler) RegisterAfterDestroyHook(afterDestroyHook func()) {
	handler.afterDestroyHooks = append(handler.afterDestroyHooks, afterDestroyHook)
}

// ExecuteInitializeHook used to execute after initialize hook
func (handler *HooksHandler) ExecuteInitializeHook(values *helpers.Dictionary) {
	if handler.initializeHook != nil {
		handler.initializeHook(values)
	}
}

// ExecuteValidationHooks used to execute validation hooks
func (handler *HooksHandler) ExecuteValidationHooks() error {
	var err error

	for _, hook := range handler.validationHooks {
		err = hook()

		if err != nil {
			break
		}
	}

	return err
}

// ExecuteAfterCreateHooks used to execute after create hooks
func (handler *HooksHandler) ExecuteAfterCreateHooks() {
	for _, hook := range handler.afterCreateHooks {
		hook()
	}
}

// ExecuteDestroyHooks used to execute after destroy hooks
func (handler *HooksHandler) ExecuteDestroyHooks() {
	for _, hook := range handler.afterDestroyHooks {
		hook()
	}
}
