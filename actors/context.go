package actors

import (
	"fmt"
	"strconv"

	"github.com/kooinam/fabio/helpers"
)

// Context used to represent context with data
type Context struct {
	params helpers.H
}

// makeContext use to instantiate controller context instance
func makeContext(params helpers.H) *Context {
	context := &Context{
		params: params,
	}

	return context
}

// ParamsStr used to retrieve params value in string
func (context *Context) ParamsStr(key string) string {
	value := context.params[key]

	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}

// Params used to retrieve params value
func (context *Context) Params(key string) interface{} {
	value := context.params[key]

	return value
}

// ParamsInt used to retrieve params value in int
func (context *Context) ParamsInt(key string, fallback int) int {
	value := context.params[key]

	if value == nil {
		return fallback
	}

	switch value.(type) {
	case string:
		i, err := strconv.Atoi(value.(string))

		if err != nil {
			return fallback
		}

		return i
	case float64:
		return int(value.(float64))
	}

	return value.(int)
}
