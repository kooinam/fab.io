package actors

// Actable is the interface for all models implementing actor model
type Actable interface {
	GetCollectionName() string
	GetID() string
	RegisterActions(*ActionsHandler)
}
