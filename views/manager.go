package views

import (
	"github.com/kooinam/fabio/helpers"
)

// Manager is singleton manager for view module
type Manager struct {
	viewHandlers map[string]*ViewHandler
}

// Setup used to setup view manager
func (manager *Manager) Setup() {
	manager.viewHandlers = make(map[string]*ViewHandler)
}

// RegisterView used to register view
func (manager *Manager) RegisterView(viewName string, handler func() Viewable) {
	viewHandler := makeViewHandler(handler)

	manager.viewHandlers[viewName] = viewHandler
}

// RenderView used to render view
func (manager *Manager) RenderView(viewName string, params helpers.H) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	renderer := viewHandler.newHandler()
	context := makeContext(manager, params)

	view := renderer.Render(context)

	return view
}

func (manager *Manager) PrepareRender(viewName string) *Renderer {
	renderer := makeRenderer(manager, viewName)

	return renderer
}
