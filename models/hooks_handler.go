package models

import (
	"github.com/kooinam/fabio/helpers"
)

// Validator is alias for func(*Connection) (interface{}, error)
type Validator func() error

// HooksHandler used to mange callbacks for models
type HooksHandler struct {
	initializeHook   func(*helpers.Dictionary)
	validationHooks  []Validator
	afterCreateHooks []func()
}

// makeHooksHandler used to instantiate CallbacksHandler
func makeHooksHandler() *HooksHandler {
	hooksHandler := &HooksHandler{
		validationHooks: []Validator{},
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

// ExecuteAfterCreateHook used to execute after create hook
func (handler *HooksHandler) ExecuteAfterCreateHook() {
	for _, hook := range handler.afterCreateHooks {
		hook()
	}
}
