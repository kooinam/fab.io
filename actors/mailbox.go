package actors

import (
	"fmt"
	"sync"

	"github.com/kooinam/fabio/helpers"
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
	actor := makeRootActor(actable)

	err = mailbox.setActor(actorIdentifier, actor)

	if err != nil {
		return err
	}

	actable.RegisterActorActions(actor.actionsHandler)
	actor.handleRegistered()

	actor.start()

	go func() {
		mailbox.manager.Tell(actorIdentifier, "Start", helpers.H{})
	}()

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
	actor := makeActor(actable, parentActor)

	err = mailbox.setActor(actorIdentifier, actor)

	if err != nil {
		return err
	}

	actable.RegisterActorActions(actor.actionsHandler)
	actor.handleRegistered()

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