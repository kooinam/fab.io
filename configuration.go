package fab

// Configuration used to configure application wide settings
type Configuration struct {
	port        string
	httpHandler func()
}

func (configuration *Configuration) SetPort(port string) {
	configuration.port = port
}

func (configuration *Configuration) SetHttpHandler(httpHandler func()) {
	configuration.httpHandler = httpHandler
}
