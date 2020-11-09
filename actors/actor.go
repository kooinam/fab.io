package actors

import (
	"sync"
	"time"
)

// Actor is the base representation of actor in actor model
type Actor struct {
	manager        *Manager
	identifier     string
	root           *Actor
	parent         *Actor
	children       []*Actor
	actionsHandler *ActionsHandler
	ch             chan event
	isRoot         bool
	messages       []*Message
	mutex          *sync.Mutex
	lastRunnedAt   time.Time
}

// makeRootActor used to instantiate root actor
func makeRootActor(manager *Manager, actable Actable) *Actor {
	identifier := actable.GetActorIdentifier()

	actor := &Actor{
		identifier: identifier,
		children:   []*Actor{},
		ch:         make(chan event),
		isRoot:     true,
		messages:   []*Message{},
		mutex:      &sync.Mutex{},
	}
	actor.actionsHandler = makeActionsHandler(manager, actor)
	actor.root = actor

	return actor
}

// makeChildActor used to instantiate actor instance
func makeChildActor(manager *Manager, actable Actable, parent *Actor) *Actor {
	root := parent.root
	identifier := actable.GetActorIdentifier()

	actor := &Actor{
		identifier: identifier,
		root:       root,
		parent:     parent,
		children:   []*Actor{},
		ch:         root.ch,
	}
	actor.actionsHandler = makeActionsHandler(manager, actor)

	return actor
}

// Identifier used to identify mailbox's address of actor
func (actor *Actor) Identifier() string {
	return actor.identifier
}

// Root used to retrieve root parent of actor
func (actor *Actor) Root() *Actor {
	return actor.root
}

// RootIdentifier used to retrieve root parent's identifier of actor
func (actor *Actor) RootIdentifier() string {
	return actor.root.Identifier()
}

// Children used to retrieve actor's children actors
func (actor *Actor) Children() []*Actor {
	return actor.children
}

// ChildrenIdentifiers used to retrieve actor's children actors' identifiers
func (actor *Actor) ChildrenIdentifiers() []string {
	identifiers := []string{}

	for _, child := range actor.children {
		identifiers = append(identifiers, child.Identifier())
	}

	return identifiers
}

// LastRunnedAt used to retrieve actor's last runned at
func (actor *Actor) LastRunnedAt() time.Time {
	return actor.lastRunnedAt
}

func (actor *Actor) pushMessage(message *Message) {
	actor.mutex.Lock()
	defer actor.mutex.Unlock()

	actor.messages = append(actor.messages, message)
}

func (actor *Actor) start() {
	actor.ch = make(chan event, 5)

	go func() {
		for event := range actor.ch {
			handled := actor.handleEvent(event)

			if !handled {
				event.nak("no registered actor for action")
			}
		}
	}()
}

func (actor *Actor) handleEvent(event event) bool {
	handled := false

	for _, child := range actor.children {
		if child.handleEvent(event) {
			handled = true
		}
	}

	if event.cascade || actor.identifier == event.actorIdentifier {
		// execute event handler
		actor.actionsHandler.handleEvent(event)

		handled = true
	}

	return handled
}

func (actor *Actor) registered() {
	parent := actor.parent

	if parent != nil {
		// is child actor
		parent.children = append(parent.children, actor)
	}
}

func (actor *Actor) unregistered() {
	parent := actor.parent

	if parent != nil {
		// is child actor
		index := -1

		for i, child := range parent.children {
			if child.identifier == actor.identifier {
				index = i

				break
			}
		}

		if index != -1 {
			parent.children = append(parent.children[:index], parent.children[index+1:]...)
		}
	} else {
		// is root actor
		close(actor.ch)
	}
}
