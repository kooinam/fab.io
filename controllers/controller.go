package controllers

// Controller used to represent base classes for all controllers
type Controller interface {
	AddBeforeActions(*CallbacksHandler)
	AddActions(*ActionsHandler)
}
