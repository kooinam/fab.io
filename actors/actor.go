package actors

// Actor is the base representation of actor in actor model
type Actor struct {
	manager         *Manager
	identifier      string
	root            *Actor
	actionsHandler  *ActionsHandler
	actionsHandlers []*ActionsHandler
	ch              chan event
}

// makeRootActor used to instantiate root actor
func makeRootActor(manager *Manager, actable Actable) *Actor {
	identifier := actable.GetActorIdentifier()

	actor := &Actor{
		identifier:      identifier,
		actionsHandler:  makeActionsHandler(manager, identifier),
		actionsHandlers: []*ActionsHandler{},
		ch:              make(chan event),
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
		actionsHandler: makeActionsHandler(manager, identifier),
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
	actor.root.actionsHandlers = append(actor.root.actionsHandlers, actor.actionsHandler)
}

func (actor *Actor) start() {
	actor.ch = make(chan event, 5)

	go func() {
		for event := range actor.ch {
			handled := false

			for _, actionsHandler := range actor.actionsHandlers {
				if actionsHandler.identifier == event.actorIdentifier {
					handled = true

					actionsHandler.handleEvent(event)
				}
			}

			if !handled {
				event.nak("no registered actor for action")
			}
		}
	}()
}
