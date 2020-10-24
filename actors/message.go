package actors

import "github.com/kooinam/fab.io/helpers"

type Message struct {
	topic  string
	params helpers.H
}

func makeMessage(topic string, params helpers.H) *Message {
	message := &Message{
		topic:  topic,
		params: params,
	}

	return message
}
