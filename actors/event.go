package actors

type Event struct {
	name       string
	parameters []string
	params     map[string]interface{}
	resCh      chan Response
}

func makeEvent(eventName string, params map[string]interface{}, resCh chan Response) *Event {
	event := &Event{
		name:   eventName,
		params: params,
		resCh:  resCh,
	}

	return event
}

// dispatch used to send event to channel without blocking. if channel's buffer is full, error response will be returned
func (event *Event) dispatch(ch chan Event) {
	select {
	case ch <- *event:
	default:
		if event.resCh != nil {
			response := makeResponse(1, "server is busy")

			event.resCh <- *response
		}
	}
}
