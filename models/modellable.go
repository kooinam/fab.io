package models

type Modellable interface {
	Initialize(*Collection)
	GetID() string
}
