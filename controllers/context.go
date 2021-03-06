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

// QueryStr used to retrieve params value in string
func (context *Context) QueryStr(key string) string {
	url := context.conn.URL()

	return url.Query().Get(key)
}

// Param used to retrieve params value
func (context *Context) Param(key string) interface{} {
	return context.params.Value(key)
}

// ParamStr used to retrieve params value in string
func (context *Context) ParamStr(key string) string {
	return context.params.ValueStr(key)
}

// ParamInt used to retrieve params value in int
func (context *Context) ParamInt(key string, fallback int) int {
	return context.params.ValueInt(key, fallback)
}

// ParamFloat64 used to retrieve params value in float64
func (context *Context) ParamFloat64(key string, fallback float64) float64 {
	return context.params.ValueFloat64(key, fallback)
}

// ParamBool used to retrieve params value in bool
func (context *Context) ParamBool(key string) bool {
	return context.params.ValueBool(key)
}

// ParamDict used to retrieve params value in dict
func (context *Context) ParamDict(key string) *helpers.Dictionary {
	return context.params.ValueDict(key)
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
