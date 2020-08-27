package models

// Modellable is the interface for all models implementing model's functionalities
type Modellable interface {
	Instantiate(*Collection, *HooksHandler, Modellable)
	Save() error
	GetID() string
	GetHooksHandler() *HooksHandler
}
