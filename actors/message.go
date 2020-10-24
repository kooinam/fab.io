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

// ParamDicts used to retrieve list of dicts
func (message *Message) ParamDicts(key string) []*helpers.Dictionary {
	return message.params.ValueDicts(key)
}
