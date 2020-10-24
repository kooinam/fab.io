package actors

import "sync"

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
	actor.root = actor

	return actor
}

// makeActor used to instantiate actor instance
func makeActor(manager *Manager, actable Actable, parent *Actor) *Actor {
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
func (actor *Actor) Root() string {
	return actor.root.Identifier()
}

func (actor *Actor) pushMessage(message *Message) {
	actor.mutex.Lock()
	defer actor.mutex.Unlock()

	actor.messages = append(actor.messages, message)
}

func (actor *Actor) handleRegistered() {
	if actor.parent != nil {
		actor.parent.children = append(actor.parent.children, actor)
	}
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
