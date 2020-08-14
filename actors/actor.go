package actors

// Actor is the base representation of actor in actor model
type Actor struct {
	actionsHandler *ActionsHandler
	Ch             chan Event
}

// makeActor used to instantiate runner instance
func makeActor(actable Actable) *Actor {
	actor := &Actor{
		actionsHandler: makeActionsHandler(),
		Ch:             make(chan Event),
	}

	actable.RegisterActions(actor.actionsHandler)

	actor.start()

	return actor
}

func (actor *Actor) start() {
	go func() {
		for event := range actor.Ch {
			actor.actionsHandler.handleEvent(event)
		}
	}()
}
