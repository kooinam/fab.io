package actors

import (
	"fmt"
	"time"

	"github.com/kooinam/fabio/helpers"
)

// Mailboxes is alias for map[string]chan Event
type Mailboxes map[string]*ActorInfo

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
	actorIdentifier := actor.Identifier()

	_, exists := manager.mailboxes[actorIdentifier]

	if exists {
		panic("actor already registered")
	}

	manager.mailboxes[actorIdentifier] = makeActorInfo(actorIdentifier, actor.ch)

	go func() {
		manager.Tell(actorIdentifier, "Start", helpers.H{})
	}()

	return actor
}

// Tell used to delegating a task to an actor asynchronously
func (manager *Manager) Tell(actorIdentifier string, eventName string, params map[string]interface{}) {
	ch := manager.mailboxes[actorIdentifier].ch
	event := makeEvent(eventName, params, nil)

	event.dispatch(ch)
}

// Request used to delegating a task to an actor synchronously with an response
func (manager *Manager) Request(actorIdentifier string, eventName string, params map[string]interface{}) error {
	var err error
	ch := manager.mailboxes[actorIdentifier].ch
	resCh := make(chan Response)
	event := makeEvent(eventName, params, resCh)

	event.dispatch(ch)
	res := <-resCh

	if res.status != 0 {
		err = fmt.Errorf(res.message)
	}

	return err
}

// GetActorInfos used to return all registered actors
func (manager *Manager) GetActorInfos() []*ActorInfo {
	actorInfos := []*ActorInfo{}

	for _, actorInfo := range manager.mailboxes {
		actorInfos = append(actorInfos, actorInfo)
	}

	return actorInfos
}

func (manager *Manager) update() {
	time.Sleep(1 * time.Second)

	for actorIdentifier := range manager.mailboxes {
		manager.Tell(actorIdentifier, "Update", helpers.H{})
	}
}
