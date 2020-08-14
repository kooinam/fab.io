package actors

type Event struct {
	name       string
	parameters []string
}

func makeEvent(eventName string) *Event {
	event := &Event{
		name: eventName,
	}

	return event
}
