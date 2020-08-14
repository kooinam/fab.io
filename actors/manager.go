package actors

// Mailboxes is alias for map[string]chan Event
type Mailboxes map[string]chan Event

// Manager is singleton manager for actor module
type Manager struct {
	mailboxes Mailboxes
}

// Setup used to setup actor manager
func (manager *Manager) Setup() {
	manager.mailboxes = make(Mailboxes)

	go func() {
		for {
			manager.run()
		}
	}()
}

// RegisterActor used to creating an actor instance for model
func (manager *Manager) RegisterActor(actable Actable) *Actor {
	actor := makeActor(actable)

	return actor
}

func (manager *Manager) run() {
	// runner.handler()

	// time.Sleep(runner.interval * time.Second)

	// runner.elapsed += runner.interval * time.Second
}
