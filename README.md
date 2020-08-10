# Fab.io

Fab.io is a lightweight game backend framework written in Go (Golang).

  - MVC Pattern
  - Powered by socket.io

### Installation

To install Fab.io package, you need to install Go and set your Go workspace first.

1. Intstall fab.io
```sh
$ go get -u github.com/kooinam/fabio
```
2. Import it in your code:

```go
import (
	fab "github.com/kooinam/fabio"
)
```

### Quick Start

1. Create an empty folder
```sh
$ mkdir fabio-chat-demo
$ cd fabio-chat-demo
```
2. Start by creating an simple Javascript chatroom application which will be connecting to our backend services.
3. Create an empty directory `demo` to hold our Javascript application codes.
```sh
$ mkdir demo
```
4. Create an HTML file `chat.html` in `demo` folder and copy the [snippet content](https://raw.githubusercontent.com/kooinam/fabio-chat-demo/master/demo/chat.html) over to `chat.html`.
5. Now proceed to setup our backend services. Use `go mod` to manage our package dependencies.
```sh
$ go mod init
```
6. Install fab.io.
```sh
$ go get -u github.com/kooinam/fabio
```
7. Create an empty directory `controllers` to hold our controllers. An controller is reponsible for handling any request and producing the appropriate output. Every controller should implement two functions `AddBeforeActions` and `AddActions`.
```sh
$ mkdir controllers
```
8. Create an go file `chat_controller.go` in `controllers` folder. Put the following snippet content into `chat_controller.go`.
```go
package controllers

import (
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/controllers"
)

// ChatController used for chat purposes
type ChatController struct {
}

// AddBeforeActions used to add before actions callbacks
func (controller *ChatController) AddBeforeActions(callbacksHandler *controllers.CallbacksHandler) {
}

// AddActions used to add actions
func (controller *ChatController) AddActions(actionsHandler *controllers.ActionsHandler) {
	actionsHandler.AddAction("Join", controller.join)
	actionsHandler.AddAction("Message", controller.message)
}

// join used for player to join a room
func (controller *ChatController) join(connection *controllers.Connection) (interface{}, error) {
	var err error
	roomID := connection.ParamsStr("roomID")

	// leave all previously joined rooms, and join new room
	connection.SingleJoin(roomID)

	return nil, err
}

// message used for player to send message message to room
func (controller *ChatController) message(connection *controllers.Connection) (interface{}, error) {
	var err error
	roomID := connection.ParamsStr("roomID")
	message := connection.ParamsStr("message")

	// broadcast message to room
	fab.BroadcastEvent("chat", roomID, "Message", nil, fab.H{
		"message": message,
	})

	return nil, err
}

```
9. Lastly, create `main.go` in root directory and put the following snippet content into `main.go`.
```go
package main

import (
	"fabio-chat-demo/controllers"
	"net/http"

	fab "github.com/kooinam/fabio"
)

func main() {
	fab.Setup()

	fab.RegisterController("chat", &controllers.ChatController{})

	fab.Serve(func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
```
10. Start our application by running
```sh
go run main.go
```
11. Navigate your browser to `http://0.0.0.0:8000` to see our chatroom application in action.

### Examples
- https://github.com/kooinam/fabio-chat-demo - An simple chatroom application with demonstrations of using controllers and routings.
- https://github.com/kooinam/fabio-demo - An simple tic-tac-toe application with demonstrations of an MVC pattern architecture.

### Dependencies
| Package | Link |
| ------ | ------ |
| go-socket.io | github.com/googollee/go-socket.io |

### Todos

 - Write MORE Tests
 - Tutorials and Documentations
 - Containerize Solutions
 - Distributed Solutions
 - Graceful Shutdown

License
----

MIT

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
