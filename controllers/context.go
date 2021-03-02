package controllers

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/views"
)

// Context used to represent context with properties and params
type Context struct {
	viewsManager *views.Manager
	conn         socketio.Conn
	properties   *helpers.Dictionary
	params       *helpers.Dictionary
	result       *Result
}

// makeContext use to instantiate controller context instance
func makeContext(viewsManager *views.Manager, conn socketio.Conn, params helpers.H) *Context {
	context := &Context{
		viewsManager: viewsManager,
		conn:         conn,
		params:       helpers.MakeDictionary(params),
		properties:   helpers.MakeDictionary(helpers.H{}),
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

// QueryStr used to retrieve params value in string
func (context *Context) QueryStr(key string) string {
	url := context.conn.URL()

	return url.Query().Get(key)
}

// ParamsInt used to retrieve params value in int
func (context *Context) ParamsInt(key string, fallback int) int {
	return context.params.ValueInt(key, fallback)
}

// ParamsFloat64 used to retrieve params value in float64
func (context *Context) ParamsFloat64(key string, fallback float64) float64 {
	return context.params.ValueFloat64(key, fallback)
}

// ParamsBool used to retrieve params value in bool
func (context *Context) ParamsBool(key string) bool {
	return context.params.ValueBool(key)
}

// SetSuccessResult used to halt controller's chain and acknowledge request with content
func (context *Context) SetSuccessResult(content interface{}) {
	context.result = makeResult()

	context.result.Set(content, StatusSuccess, nil)
}

// SetErrorResult used to halt controller's chain and acknowledge request with error status and error
func (context *Context) SetErrorResult(err error) {
	context.result = makeResult()

	context.result.Set(nil, StatusError, err)
}

// SetUnauthorizedResult used to halt controller's chain and acknowledge request with unthorized status
func (context *Context) SetUnauthorizedResult() {
	context.result = makeResult()

	context.result.Set(nil, StatusError, fmt.Errorf("unauthorized"))
}

func (context *Context) PrepareRender(viewName string) *views.Renderer {
	return context.viewsManager.PrepareRender(viewName)
}
