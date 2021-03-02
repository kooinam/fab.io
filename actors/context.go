package actors

import (
	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/views"
)

// Context used to represent actor execution context with data
type Context struct {
	viewsManager *views.Manager
	actor        *Actor
	params       *helpers.Dictionary
	properties   *helpers.Dictionary
	result       string
}

// makeContext use to instantiate controller context instance
func makeContext(viewsManager *views.Manager, actor *Actor, params helpers.H) *Context {
	context := &Context{
		viewsManager: viewsManager,
		actor:        actor,
		params:       helpers.MakeDictionary(params),
		properties:   helpers.MakeDictionary(helpers.H{}),
	}

	return context
}

func (context *Context) ParamsDict() *helpers.Dictionary {
	return context.params
}

// Params used to retrieve params value
func (context *Context) Params(key string) interface{} {
	return context.params.Value(key)
}

// ParamsStr used to retrieve params value in string
func (context *Context) ParamsStr(key string) string {
	return context.params.ValueStr(key)
}

// ParamsInt used to retrieve params value in int
func (context *Context) ParamsInt(key string, fallback int) int {
	return context.params.ValueInt(key, fallback)
}

// ParamsFloat64 used to retrieve params value in float64
func (context *Context) ParamsFloat64(key string, fallback float64) float64 {
	return context.params.ValueFloat64(key, fallback)
}

func (context *Context) ParamDicts(key string) []*helpers.Dictionary {
	return context.params.ValueDicts(key)
}

// ParamsBool used to retrieve params value in bool
func (context *Context) ParamsBool(key string) bool {
	return context.params.ValueBool(key)
}

// SetProperty used to set property that can used across the context
func (context *Context) SetProperty(key string, value interface{}) {
	context.properties.Set(key, value)
}

// Property used to retrieve context property value
func (context *Context) Property(key string) interface{} {
	return context.properties.Value(key)
}

// PropertyStr used to retrieve params value in string
func (context *Context) PropertyStr(key string) string {
	return context.properties.ValueStr(key)
}

// PropertyInt used to retrieve params value in int
func (context *Context) PropertyInt(key string, fallback int) int {
	return context.properties.ValueInt(key, fallback)
}

// PropertyFloat64 used to retrieve params value in float64
func (context *Context) PropertyFloat64(key string, fallback float64) float64 {
	return context.properties.ValueFloat64(key, fallback)
}

// PropertyBool used to retrieve params value in bool
func (context *Context) PropertyBool(key string) bool {
	return context.properties.ValueBool(key)
}

func (context *Context) PrepareRender(viewName string) *views.Renderer {
	return context.viewsManager.PrepareRender(viewName)
}

// Tell used to delegating a task to an actor asynchronously
func (context *Context) Tell(actor *Actor, eventName string, params map[string]interface{}, cascade bool) {
	ch := actor.ch
	event := makeEvent(actor.Identifier(), eventName, params, nil, cascade)

	event.dispatch(ch)
}

// PopMessage used to pop message from messages queue, return nil if queue is empty
func (context *Context) PopMessage() *Message {
	actor := context.actor

	actor.mutex.Lock()
	defer actor.mutex.Unlock()

	var message *Message

	if len(actor.messages) > 0 {
		message = actor.messages[0]

		actor.messages = actor.messages[1:]
	}

	return message
}
