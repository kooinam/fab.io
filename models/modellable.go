package models

// Modellable is the interface for all models implementing model's functionalities
type Modellable interface {
	GetID() string
	Save() error
	Memoize()

	_initialize(*Collection, *HooksHandler, Modellable)
	getHooksHandler() *HooksHandler
}
