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
func (manager *Manager) RenderView(viewName string, properties helpers.H, rootKey string) interface{} {
	v := manager.RenderViewWithoutRootKey(viewName, properties)

	return helpers.BuildJSON(v, true, rootKey)
}

// RenderViewWithoutRootKey used to render view without root key
func (manager *Manager) RenderViewWithoutRootKey(viewName string, properties helpers.H) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	dict := helpers.MakeDictionary(properties)

	view := viewHandler.newHandler()
	v := view.Render(dict)

	return v
}

// RenderViews used to render collection of views
func (manager *Manager) RenderViews(viewName string, propertiesList []helpers.H, rootKey string) interface{} {
	vs := manager.RenderViewsWithoutRootKey(viewName, propertiesList)

	return helpers.BuildJSON(vs, true, rootKey)
}

// RenderViewsWithoutRootKey used to render collection of views without root key
func (manager *Manager) RenderViewsWithoutRootKey(viewName string, propertiesList []helpers.H) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	vs := make([]interface{}, len(propertiesList))

	for i, properties := range propertiesList {
		dict := helpers.MakeDictionary(properties)

		view := viewHandler.newHandler()
		v := view.Render(dict)

		vs[i] = v
	}

	return vs
}
