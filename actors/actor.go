package actors

import (
	"fmt"
)

// Actor is the base representation of actor in actor model
type Actor struct {
	identifier     string
	actionsHandler *ActionsHandler
	ch             chan Event
}

// makeActor used to instantiate runner instance
func makeActor(actable Actable) *Actor {
	actor := &Actor{
		identifier:     fmt.Sprintf(actable.GetActorIdentifier()),
		actionsHandler: makeActionsHandler(),
		ch:             make(chan Event, 5),
	}

	actable.RegisterActorActions(actor.actionsHandler)

	actor.start()

	return actor
}

// Identifier used to identify mailbox's address of actor
func (actor *Actor) Identifier() string {
	return actor.identifier
}

func (actor *Actor) start() {
	go func() {
		for event := range actor.ch {
			actor.actionsHandler.handleEvent(event)
		}
	}()
}
