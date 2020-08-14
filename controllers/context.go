package controllers

import (
	"fmt"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
)

// Context used to represent context with properties and params
type Context struct {
	conn       socketio.Conn
	properties map[string]interface{}
	params     map[string]interface{}
}

// makeConnection use to instantiate connection instance
func makeContext(conn socketio.Conn, params map[string]interface{}) *Context {
	context := &Context{
		conn:       conn,
		params:     params,
		properties: make(map[string]interface{}),
	}

	return context
}

// Join used to join socketio room
func (context *Context) Join(room string) {
	context.conn.Join(room)
}

// SingleJoin used to join socketio room while leaving other joined rooms
func (context *Context) SingleJoin(room string) {
	for _, joinedRoom := range context.Rooms() {
		if joinedRoom != room {
			context.Leave(joinedRoom)
		}
	}

	context.Join(room)
}

// Leave used to leave socketio room
func (context *Context) Leave(room string) {
	context.conn.Leave(room)
}

// Rooms used to retrieve all connection's rooms
func (context *Context) Rooms() []string {
	return context.conn.Rooms()
}

// SetProperty used to set property that can used across the context
func (context *Context) SetProperty(key string, value interface{}) {
	context.properties[key] = value
}

// Property used to retrieve context property value
func (context *Context) Property(key string) interface{} {
	value := context.properties[key]

	return value
}

// PropertyWithFallback used to retrieve context property value with a fallback
func (context *Context) PropertyWithFallback(key string, fallback interface{}) interface{} {
	value := context.properties[key]

	if value != nil {
		return value
	}

	return fallback
}

// Params used to retrieve params value
func (context *Context) Params(key string) interface{} {
	value := context.params[key]

	return value
}

// ParamsStr used to retrieve params value in string
func (context *Context) ParamsStr(key string) string {
	value := context.params[key]

	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
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

// ParamsWithFallback used to retrieve params value with a fallback
func (context *Context) ParamsWithFallback(key string, fallback interface{}) interface{} {
	value := context.params[key]

	if value != nil {
		return value
	}

	return fallback
}
