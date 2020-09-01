package actors

import (
	"fmt"
	"sync"
	"time"

	"github.com/kooinam/fabio/helpers"
)

// Mailboxes is alias for map[string]chan Event
type Mailboxes map[string]*ActorInfo

// Manager is singleton manager for actor module
type Manager struct {
	mailboxes Mailboxes
	mutex     *sync.RWMutex
}

// Setup used to setup actor manager
func (manager *Manager) Setup() {
	manager.mailboxes = make(Mailboxes)
	manager.mutex = &sync.RWMutex{}

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

// RegisterActor used to creating an actor instance for model
func (manager *Manager) RegisterActor(nsp string, actable Actable) *Actor {
	actor := makeActor(nsp, actable)
	actorIdentifier := actor.Identifier()

	_, exists := manager.mailboxes[actorIdentifier]

	if exists {
		panic("actor already registered")
	}

	actorInfo := makeActorInfo(actorIdentifier, actor.ch)
	manager.setActorInfo(actorIdentifier, actorInfo)

	go func() {
		manager.Tell(actorIdentifier, "Start", helpers.H{})
	}()

	return actor
}

// Tell used to delegating a task to an actor asynchronously
func (manager *Manager) Tell(actorIdentifier string, eventName string, params map[string]interface{}) {
	actorInfo := manager.getActorInfo(actorIdentifier)
	ch := actorInfo.ch
	event := makeEvent(eventName, params, nil)

	event.dispatch(ch)
}

// Request used to delegating a task to an actor synchronously with an response
func (manager *Manager) Request(actorIdentifier string, eventName string, params map[string]interface{}) error {
	var err error

	actorInfo := manager.getActorInfo(actorIdentifier)
	ch := actorInfo.ch
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
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	actorInfos := []*ActorInfo{}

	for _, actorInfo := range manager.mailboxes {
		actorInfos = append(actorInfos, actorInfo)
	}

	return actorInfos
}

func (manager *Manager) update(dt float64) {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	for actorIdentifier := range manager.mailboxes {
		manager.Tell(actorIdentifier, "Update", helpers.H{
			"dt": dt,
		})
	}
}

func (manager *Manager) getActorInfo(actorIdentifier string) *ActorInfo {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	return manager.mailboxes[actorIdentifier]
}

func (manager *Manager) setActorInfo(actorIdentifier string, actorInfo *ActorInfo) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	manager.mailboxes[actorIdentifier] = actorInfo
}
