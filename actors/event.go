package actors

type event struct {
	actorIdentifier string
	name            string
	params          map[string]interface{}
	resCh           chan Response
	cascade         bool
}

func makeEvent(actorIdentifier string, eventName string, params map[string]interface{}, resCh chan Response, cascade bool) *event {
	event := &event{
		actorIdentifier: actorIdentifier,
		name:            eventName,
		params:          params,
		resCh:           resCh,
		cascade:         cascade,
	}

	return event
}

// dispatch used to send event to channel without blocking. if channel's buffer is full, error response will be returned
func (event *event) dispatch(ch chan event) {
	select {
	case ch <- *event:
	default:
		event.nak("server is busy")
	}
}

func (event *event) ack() {
	if event.resCh != nil {
		response := makeResponse(0, "ok")

		event.resCh <- *response
	}
}

func (event *event) nak(message string) {
	if event.resCh != nil {
		response := makeResponse(1, message)

		event.resCh <- *response
	}
}
