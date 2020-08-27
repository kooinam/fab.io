package models

import (
	"github.com/kooinam/fabio/helpers"
)

// HooksHandler used to mange callbacks for models
type HooksHandler struct {
	afterInitializeHook func(*helpers.Dictionary)
	validationHooks     []func() error
}

// makeHooksHandler used to instantiate CallbacksHandler
func makeHooksHandler() *HooksHandler {
	hooksHandler := &HooksHandler{
		validationHooks: []func() error{},
	}

	return hooksHandler
}

// RegisterAfterInitializeHook used to add after initialize hook
func (handler *HooksHandler) RegisterAfterInitializeHook(afterInitializeHook func(*helpers.Dictionary)) {
	handler.afterInitializeHook = afterInitializeHook
}

// RegisterValidationHook used to add a validation hook
func (handler *HooksHandler) RegisterValidationHook(validationHook func() error) {
	handler.validationHooks = append(handler.validationHooks, validationHook)
}

// executeAfterInitializeHook used to execute after initialize hook
func (handler *HooksHandler) executeAfterInitializeHook(values *helpers.Dictionary) {
	if handler.afterInitializeHook != nil {
		handler.afterInitializeHook(values)
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
