package actors

import (
	"strings"

	"github.com/kooinam/fab.io/logger"
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
}

// makeRootActor used to instantiate root actor
func makeRootActor(manager *Manager, actable Actable) *Actor {
	identifier := actable.GetActorIdentifier()

	actor := &Actor{
		identifier:     identifier,
		children:       []*Actor{},
		actionsHandler: makeActionsHandler(manager),
		ch:             make(chan event),
		isRoot:         true,
	}

	actor.root = actor

	return actor
}

// makeActor used to instantiate actor instance
func makeActor(manager *Manager, actable Actable, parent *Actor) *Actor {
	root := parent.root
	identifier := actable.GetActorIdentifier()

	actor := &Actor{
		identifier:     identifier,
		root:           root,
		parent:         parent,
		children:       []*Actor{},
		actionsHandler: makeActionsHandler(manager),
		ch:             root.ch,
	}

	return actor
}

// Identifier used to identify mailbox's address of actor
func (actor *Actor) Identifier() string {
	return actor.identifier
}

func (actor *Actor) Ch() chan event {
	return actor.ch
}

func (actor *Actor) Root() string {
	return actor.root.Identifier()
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
			if strings.Contains(actor.Identifier(), "player") {
				logger.Debug("start...")
			}
			handled := actor.handleEvent(event)
			if strings.Contains(actor.Identifier(), "player") {
				logger.Debug("end...")
			}

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
		actor.actionsHandler.handleEvent(event)

		handled = true
	}

	return handled
}
