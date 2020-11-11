package actors

import "github.com/kooinam/fab.io/helpers"

type Message struct {
	topic  string
	params *helpers.Dictionary
}

func makeMessage(topic string, params helpers.H) *Message {
	message := &Message{
		topic:  topic,
		params: helpers.MakeDictionary(params),
	}

	return message
}

// Topic used to retrieve message's topic
func (message *Message) Topic() string {
	return message.topic
}

// ParamDicts used to retrieve list of dicts
func (message *Message) ParamDicts(key string) []*helpers.Dictionary {
	return message.params.ValueDicts(key)
}

// ParamInt used to retrieve int
func (message *Message) ParamInt(key string, fallback int) int {
	return message.params.ValueInt(key, fallback)
}

func (message *Message) Params(key string) *helpers.Dictionary {
	return message.params
}
