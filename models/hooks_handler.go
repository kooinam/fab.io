package models

import (
	"github.com/kooinam/fabio/helpers"
)

// HooksHandler used to mange callbacks for models
type HooksHandler struct {
	initializeHook       func(*helpers.Dictionary)
	afterInstantiateHook func()
	validationHooks      []func() error
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

// RegisterAfterInstantiateHook used to add after instantiate hook
func (handler *HooksHandler) RegisterAfterInstantiateHook(afterInstantiateHook func()) {
	handler.afterInstantiateHook = afterInstantiateHook
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

// executeAfterInstantiateHook used to execute after instantiate hook
func (handler *HooksHandler) executeAfterInstantiateHook() {
	if handler.afterInstantiateHook != nil {
		handler.afterInstantiateHook()
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
