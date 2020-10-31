package actors

import (
	"fmt"
	"sync"

	"github.com/kooinam/fab.io/helpers"
)

// Addresses is alias for map[string]*Actor
type Addresses map[string]*Actor

type Mailbox struct {
	manager   *Manager
	mutex     *sync.RWMutex
	addresses Addresses
}

func makeMailbox(manager *Manager) *Mailbox {
	mailbox := &Mailbox{
		manager:   manager,
		mutex:     &sync.RWMutex{},
		addresses: make(Addresses),
	}

	return mailbox
}

// registerActor used for registering an actor
func (mailbox *Mailbox) registerActor(actable Actable) error {
	var err error

	actorIdentifier := actable.GetActorIdentifier()
	actor := makeRootActor(mailbox.manager, actable)

	err = mailbox.setActor(actorIdentifier, actor)

	if err != nil {
		return err
	}

	actable.RegisterActorActions(actor.actionsHandler)
	actor.registered()

	actor.start()

	go func() {
		mailbox.manager.Tell(actorIdentifier, "Start", helpers.H{})
	}()

	return err
}

// deregisterActor used for deregistering an actor
func (mailbox *Mailbox) deregisterActor(actorIdentifier string) error {
	var err error

	actor := mailbox.getActor(actorIdentifier)
	actor.unregistered()

	mailbox.unsetActor(actorIdentifier)

	return err
}

// registerChildActor used for registering an child actor
func (mailbox *Mailbox) registerChildActor(actable Actable, parent Actable) error {
	var err error

	parentActor := mailbox.getActor(parent.GetActorIdentifier())

	if parentActor == nil {
		err = fmt.Errorf("parent actor not found")

		return err
	}

	actorIdentifier := actable.GetActorIdentifier()
	actor := makeActor(mailbox.manager, actable, parentActor)

	err = mailbox.setActor(actorIdentifier, actor)

	if err != nil {
		return err
	}

	actable.RegisterActorActions(actor.actionsHandler)
	actor.registered()

	go func() {
		mailbox.manager.Tell(actorIdentifier, "Start", helpers.H{})
	}()

	return err
}

func (mailbox *Mailbox) getActor(actorIdentifier string) *Actor {
	mailbox.mutex.RLock()
	defer mailbox.mutex.RUnlock()

	return mailbox.addresses[actorIdentifier]
}

func (mailbox *Mailbox) setActor(actorIdentifier string, actor *Actor) error {
	var err error

	mailbox.mutex.Lock()
	defer mailbox.mutex.Unlock()

	if len(actorIdentifier) == 0 {
		err = fmt.Errorf("actor identifier is empty")

		return err
	} else if mailbox.addresses[actorIdentifier] != nil {
		err = fmt.Errorf("actor already registered")

		return err
	}

	mailbox.addresses[actorIdentifier] = actor

	return err
}

func (mailbox *Mailbox) unsetActor(actorIdentifier string) error {
	var err error

	mailbox.mutex.Lock()
	defer mailbox.mutex.Unlock()

	if len(actorIdentifier) == 0 {
		err = fmt.Errorf("actor identifier is empty")

		return err
	} else if mailbox.addresses[actorIdentifier] == nil {
		err = fmt.Errorf("actor is not registered")

		return err
	}

	delete(mailbox.addresses, actorIdentifier)

	return err
}
