package controllers

// Controllable is the interface for all controllers implementing RegisterBeforeHooks and RegisterActions
type Controllable interface {
	RegisterBeforeHooks(*HooksHandler)
	RegisterActions(*ActionsHandler)
}
