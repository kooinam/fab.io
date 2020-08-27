package models

import (
	"github.com/kooinam/fabio/helpers"
)

// HooksHandler used to mange callbacks for models
type HooksHandler struct {
	initializeHook  func(*helpers.Dictionary)
	validationHooks []func() error
}

// makeHooksHandler used to instantiate CallbacksHandler
func makeHooksHandler() *HooksHandler {
	hooksHandler := &HooksHandler{
		validationHooks: []func() error{},
	}

	return hooksHandler
}

// RegisterInitializeHook used to add initialize hook
func (handler *HooksHandler) RegisterInitializeHook(initializeHook func(*helpers.Dictionary)) {
	handler.initializeHook = initializeHook
}

// RegisterValidationHook used to add a validation hook
func (handler *HooksHandler) RegisterValidationHook(validationHook func() error) {
	handler.validationHooks = append(handler.validationHooks, validationHook)
}

// executeInitializeHook used to execute after initialize hook
func (handler *HooksHandler) executeInitializeHook(values *helpers.Dictionary) {
	if handler.initializeHook != nil {
		handler.initializeHook(values)
	}
}

// executeValidationHooks used to execute validation hooks
func (handler *HooksHandler) executeValidationHooks() error {
	var err error

	for _, hook := range handler.validationHooks {
		err = hook()

		if err != nil {
			break
		}
	}

	return err
}
