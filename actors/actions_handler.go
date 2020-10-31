package actors

import (
	"time"

	"github.com/kooinam/fab.io/logger"
)

// Hook is alias for func(string, *Context)
type Hook func(string, *Context)

// Action is alias for func(context *Context) error
type Action func(context *Context) error

// ActionsHandler used to mange callbacks for controllers
type ActionsHandler struct {
	manager           *Manager
	actor             *Actor
	beforeActionHooks []Hook
	afterActionHooks  []Hook
	actions           map[string]Action
}

// makeActionsHandler used to instantiate EventsHandler
func makeActionsHandler(manager *Manager, actor *Actor) *ActionsHandler {
	actionsHandler := &ActionsHandler{
		manager: manager,
		actor:   actor,
		actions: make(map[string]Action),
	}

	return actionsHandler
}

// RegisterBeforeActionHook used to add before hook
func (handler *ActionsHandler) RegisterBeforeActionHook(beforeActionHook Hook) {
	logger.Debug("%v", handler)
	handler.beforeActionHooks = append(handler.beforeActionHooks, beforeActionHook)
}

// RegisterAfterActionHook used to add after hook
func (handler *ActionsHandler) RegisterAfterActionHook(afterActionHook Hook) {
	handler.afterActionHooks = append(handler.afterActionHooks, afterActionHook)
}

// RegisterAction used to register action
func (handler *ActionsHandler) RegisterAction(actionName string, action Action) {
	if handler.actions[actionName] != nil {
		panic("actor action registered")
	}

	handler.actions[actionName] = action
}

// executeBeforeActionHooks used to execute before action hooks
func (handler *ActionsHandler) executeBeforeActionHooks(action string, context *Context) {
	for _, hook := range handler.beforeActionHooks {
		hook(action, context)
	}
}

// executeAfterActionHooks used to execute after action hooks
func (handler *ActionsHandler) executeAfterActionHooks(action string, context *Context) {
	for _, hook := range handler.afterActionHooks {
		hook(action, context)
	}
}

func (handler *ActionsHandler) handleEvent(event event) {
	action := handler.actions[event.name]

	if action != nil {
		context := makeContext(handler.manager.viewsManager, handler.actor, event.params)

		handler.executeBeforeActionHooks(event.name, context)

		err := action(context)

		if err == nil {
			handler.actor.lastRunnedAt = time.Now()

			event.ack()

			handler.executeAfterActionHooks(event.name, context)
		} else {
			event.nak(err.Error())
		}
	} else {
		event.nak("no action found")
	}
}
