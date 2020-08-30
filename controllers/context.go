package controllers

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fabio/helpers"
)

// Context used to represent context with properties and params
type Context struct {
	conn       socketio.Conn
	properties *helpers.Dictionary
	params     *helpers.Dictionary
	result     *Result
}

// makeContext use to instantiate controller context instance
func makeContext(conn socketio.Conn, params helpers.H) *Context {
	context := &Context{
		conn:       conn,
		params:     helpers.MakeDictionary(params),
		properties: helpers.MakeDictionary(helpers.H{}),
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
	context.properties.Set(key, value)
}

// Property used to retrieve context property value
func (context *Context) Property(key string) interface{} {
	return context.properties.Value(key)
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

// SetResult used to half controller's chain and acknoledge request with content, status and error
func (context *Context) SetResult(content interface{}, status string, err error) {
	result := makeResult()

	result.Set(content, status, err)
}
