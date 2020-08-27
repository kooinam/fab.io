package models

// Modellable is the interface for all models implementing model's functionalities
type Modellable interface {
	Save() error
	GetID() string

	initialize(*Collection, *HooksHandler, Modellable)
	getHooksHandler() *HooksHandler
}
