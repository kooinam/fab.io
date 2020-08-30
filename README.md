<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/kooinam/fabio.svg?style=flat-square
[contributors-url]: https://github.com/kooinam/fabio/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/kooinam/fabio.svg?style=flat-square
[forks-url]: https://github.com/kooinam/fabio/network/members
[stars-shield]: https://img.shields.io/github/stars/kooinam/fabio.svg?style=flat-square
[stars-url]: https://github.com/kooinam/fabio/stargazers
[issues-shield]: https://img.shields.io/github/issues/kooinam/fabio.svg?style=flat-square
[issues-url]: https://github.com/kooinam/fabio/issues
[license-shield]: https://img.shields.io/github/license/kooinam/fabio.svg?style=flat-square
[license-url]: https://github.com/kooinam/fabio/blob/master/LICENSE.txt

# Fab.io
Fab.io is a lightweight real-time game backend framework written in Go (Golang).

  - [MVC Pattern](https://github.com/kooinam/fabio/wiki)
  - [Synchronized Loop Based Actor Model](https://github.com/kooinam/fabio/wiki/Actor-Model)
  - Powered by socket.io

## Table of Contents
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Resources](#resources)
* [Dependencies](#dependencies)
* [Roadmap](#roadmap)
* [License](#license)

## Getting Started

### Prerequisites

To install Fab.io package, you will need to:
1. Have Go installed. Head over to Go's download page [here](https://golang.org/dl/) to install it.
2. Setup your Go workspace.

### Installation

1. Install Fab.io
```sh
$ go get -u github.com/kooinam/fabio
```
2. Import it in your code:
```go
import (
	fab "github.com/kooinam/fabio"
)
```

## Usage
### Example - (Simple JavaScript Chatroom App):
In our first example, we start by creating an simple JavaScript chatroom application which will be connecting to backend services using Fab.io.
#### Setting up your workspace:
1. Create an empty directory. In this example we shall use `fabio-chat-demo`:
```sh
$ mkdir fabio-chat-demo
$ cd fabio-chat-demo
```
2. Create an empty directory `demo` inside `fabio-chat-demo` to hold our Javascript application codes.
```sh
$ mkdir demo
```
3. Create an HTML file `chat.html` in the `demo` folder and copy the [snippet content](https://raw.githubusercontent.com/kooinam/fabio-chat-demo/master/demo/chat.html) over into `chat.html`.

#### Setting up backend services:
Now, let's proceed to setup our backend services.
1. Use the `go mod` command to manage our package dependencies. Let's go ahead and initialize our package dependencies:
```sh
$ go mod init fabio-chat-demo
```
2. Install Fab.io.
```sh
$ go get -u github.com/kooinam/fabio
```
3. Create an empty directory `controllers` inside `fabio-chat-demo` to hold our controllers. A controller is responsible for handling any request and producing the appropriate output. Every controller should implement two functions `RegisterBeforeHooks` and `RegisterActions`.
```sh
$ mkdir controllers
```
4. Create an go file `chat_controller.go` in `controllers` folder. Put the following snippet content into `chat_controller.go`.
```go
package controllers

import (
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/controllers"
	"github.com/kooinam/fabio/helpers"
)

// ChatController used for chat purposes
type ChatController struct {
}

// RegisterHooksAndActions used to register hooks and actions
func (controller *ChatController) RegisterHooksAndActions(hooksHandler *controllers.HooksHandler, actionsHandler *controllers.ActionsHandler) {
	actionsHandler.RegisterAction("Join", controller.join)
	actionsHandler.RegisterAction("Message", controller.message)
}

// join used for player to join a room
func (controller *ChatController) join(context *controllers.Context) {
	roomID := context.ParamsStr("roomID")

	// leave all previously joined rooms, and join new room
	context.SingleJoin(roomID)
}

// message used for player to send message message to room
func (controller *ChatController) message(context *controllers.Context) {
	roomID := context.ParamsStr("roomID")
	message := context.ParamsStr("message")

	// broadcast message to room
	fab.ControllerManager().BroadcastEvent("chat", roomID, "Message", nil, helpers.H{
		"message": message,
	})
}
```
5. Lastly, create `main.go` in root directory and put the following snippet content into `main.go`.
```go
package main

import (
	"fabio-chat-demo/controllers"
	"net/http"

	fab "github.com/kooinam/fabio"
)

func main() {
	fab.Setup()

	fab.ControllerManager().RegisterController("chat", &controllers.ChatController{})

	fab.ControllerManager().Serve("8000", func() {
		fs := http.FileServer(http.Dir("./demo"))
		http.Handle("/demo/", http.StripPrefix("/demo/", fs))
	})
}
```

#### You are done!
Congrats! Now all that's left to do is run the app!
1. Start our application by running:
```sh
go run main.go
```
2. Navigate to `http://localhost:8000/demo/chat.html` on your browser to see your chatroom application in action!

#### Interested on other use cases?
Expore more example use cases by reading our [Wiki](https://github.com/kooinam/fabio/wiki)!

## Resources
- [Wiki](https://github.com/kooinam/fabio/wiki)
- https://github.com/kooinam/fabio-chat-demo - An simple chatroom demo with demonstrations of using routings and controllers.
- https://github.com/kooinam/fabio-demo - An simple tic-tac-toe demo with demonstrations of an MVC pattern architecture.

## Dependencies
| Package | Link |
| ------ | ------ |
| go-socket.io | github.com/googollee/go-socket.io |

## Roadmap
Some of our upcoming key feature(s)/improvement(s) include:
 - Write MORE Tests
 - Tutorials and Documentations
 - Containerize Solutions
 - Distributed Solutions
 - Graceful Shutdown
 - Actor Model

## License

Distributed under the MIT License.

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
