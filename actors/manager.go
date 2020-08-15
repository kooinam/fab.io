package actors

import (
	"time"
)

// Mailboxes is alias for map[string]chan Event
type Mailboxes map[string]chan Event

// Manager is singleton manager for actor module
type Manager struct {
	mailboxes Mailboxes
}

// Setup used to setup actor manager
func (manager *Manager) Setup() {
	manager.mailboxes = make(Mailboxes)

	go func() {
		for {
			manager.update()
		}
	}()
}

// RegisterActor used to creating an actor instance for model
func (manager *Manager) RegisterActor(nsp string, actable Actable) *Actor {
	actor := makeActor(nsp, actable)

	manager.mailboxes[actor.Identifier()] = actor.ch

	return actor
}

func (manager *Manager) Broadcast(actorIdentifier string, eventName string) {
	ch := manager.mailboxes[actorIdentifier]
	event := makeEvent(eventName)

	ch <- *event
}

func (manager *Manager) update() {
	time.Sleep(1 * time.Second)

	for actorIdentifier := range manager.mailboxes {
		manager.Broadcast(actorIdentifier, "Update")
	}
}
