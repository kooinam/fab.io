package models

import "github.com/kooinam/fabio/helpers"

type Context struct {
	collection          *Collection
	hooksHandler        *HooksHandler
	associationsHandler *AssociationsHandler
	attributes          *helpers.Dictionary
	item                Modellable
}

func makeContext(collection *Collection, attributes *helpers.Dictionary) *Context {
	context := &Context{
		collection:          collection,
		hooksHandler:        makeHooksHandler(),
		associationsHandler: makeAssociationsHandler(),
		attributes:          attributes,
	}

	return context
}

func (context *Context) Collection() *Collection {
	return context.collection
}

func (context *Context) SetItem(item Modellable) {
	context.item = item
}

func (context *Context) Item() Modellable {
	return context.item
}

func (context *Context) HooksHandler() *HooksHandler {
	return context.hooksHandler
}

func (context *Context) AssociationsHandler() *AssociationsHandler {
	return context.associationsHandler
}

func (context *Context) Attributes() *helpers.Dictionary {
	return context.attributes
}
