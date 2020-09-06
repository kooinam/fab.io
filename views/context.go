package views

import (
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

// Context used to represent view rendering context with data
type Context struct {
	item   models.Modellable
	params *helpers.Dictionary
}

// makeContext use to instantiate controller context instance
func makeContext(params helpers.H) *Context {
	context := &Context{
		params: helpers.MakeDictionary(params),
	}

	return context
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

// Item used to retrieve itme in context
func (context *Context) Item() models.Modellable {
	return context.item
}

func (context *Context) setItem(item models.Modellable) {
	context.item = item
}
