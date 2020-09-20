package actors

import (
	"fmt"
	"time"

	"github.com/kooinam/fab.io/helpers"
)

// Manager is singleton manager for actor module
type Manager struct {
	*Mailbox
}

// Setup used to setup actor manager
func (manager *Manager) Setup() {
	manager.Mailbox = makeMailbox(manager)

	go func() {
		t1 := time.Now()

		for {
			time.Sleep(1 * time.Second)

			t2 := time.Now()
			dt := t2.Sub(t1)
			t1 = t2

			manager.update(dt.Seconds())
		}
	}()
}

// RegisterActor used for registering an actor
func (manager *Manager) RegisterActor(actable Actable) error {
	err := manager.registerActor(actable)

	return err
}

// RegisterChildActor used to creating an actor instance for model
func (manager *Manager) RegisterChildActor(parent Actable, actable Actable) error {
	var err error

	err = manager.registerChildActor(actable, parent)

	return err
}

// Tell used to delegating a task to an actor asynchronously
func (manager *Manager) Tell(actorIdentifier string, eventName string, params map[string]interface{}) {
	actor := manager.getActor(actorIdentifier)

	if actor == nil {
		panic(fmt.Sprintf("%v not registered", actorIdentifier))
	}

	ch := actor.ch
	event := makeEvent(actorIdentifier, eventName, params, nil)

	event.dispatch(ch)
}

// Request used to delegating a task to an actor synchronously with an response
func (manager *Manager) Request(actorIdentifier string, eventName string, params map[string]interface{}) error {
	var err error

	child := manager.getActor(actorIdentifier)
	ch := child.ch
	resCh := make(chan Response)
	event := makeEvent(actorIdentifier, eventName, params, resCh)

	event.dispatch(ch)
	res := <-resCh

	if res.status != 0 {
		err = fmt.Errorf(res.message)
	}

	return err
}

// GetActors used to return all registered actors
func (manager *Manager) GetActors() []*Actor {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	actors := []*Actor{}

	for _, actor := range manager.addresses {
		actors = append(actors, actor)
	}

	return actors
}

func (manager *Manager) update(dt float64) {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	for actorIdentifier := range manager.addresses {
		manager.Tell(actorIdentifier, "Update", helpers.H{
			"dt": dt,
		})
	}
}
