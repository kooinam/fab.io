package controllers

import (
	socketio "github.com/googollee/go-socket.io"
)

// Connection used to represent connection with properties and params
type Connection struct {
	conn       socketio.Conn
	properties map[string]interface{}
	params     map[string]interface{}
}

// MakeConnection use to instantiate connection instance
func MakeConnection(conn socketio.Conn, params map[string]interface{}) *Connection {
	connection := &Connection{
		conn:       conn,
		params:     params,
		properties: make(map[string]interface{}),
	}

	return connection
}

// Join used to join socketio room
func (connection *Connection) Join(room string) {
	connection.conn.Join(room)
}

// SetProperty used to set property that can used across the context
func (connection *Connection) SetProperty(key string, value interface{}) {
	connection.properties[key] = value
}

// Property used to retrieve context property value
func (connection *Connection) Property(key string) interface{} {
	value := connection.properties[key]

	return value
}

// PropertyWithFallback used to retrieve context property value with a fallback
func (connection *Connection) PropertyWithFallback(key string, fallback interface{}) interface{} {
	value := connection.properties[key]

	if value != nil {
		return value
	}

	return fallback
}

// Params used to retrieve params value
func (connection *Connection) Params(key string, fallback interface{}) interface{} {
	value := connection.properties[key]

	return value
}

// ParamsWithFallback used to retrieve params value with a fallback
func (connection *Connection) ParamsWithFallback(key string, fallback interface{}) interface{} {
	value := connection.properties[key]

	if value != nil {
		return value
	}

	return fallback
}
