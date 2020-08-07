package controllers

import (
	socketio "github.com/googollee/go-socket.io"
)

type Connection struct {
	conn       socketio.Conn
	properties map[string]interface{}
	Params     map[string]interface{}
}

func MakeConnection(conn socketio.Conn, params map[string]interface{}) *Connection {
	connection := &Connection{
		conn:       conn,
		Params:     params,
		properties: make(map[string]interface{}),
	}

	return connection
}

// Join used to join socketio room
func (connection *Connection) Join(room string) {
	connection.conn.Join(room)
}

// GetQueryParams used to get query params
func (connection *Connection) GetQueryParams(query string) string {
	url := connection.conn.URL()

	return url.Query().Get(query)
}

func (connection *Connection) AddProperty(key string, value interface{}) {
	connection.properties[key] = value
}

func (connection *Connection) GetProperty(key string) interface{} {
	return connection.properties[key]
}
