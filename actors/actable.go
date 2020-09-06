package actors

// Actable is the interface for all models implementing actor model
type Actable interface {
	GetActorIdentifier() string
	RegisterActorActions(*ActionsHandler)
}
