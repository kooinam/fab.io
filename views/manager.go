package views

import (
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
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
	context := makeContext(params)

	view := renderer.Render(context)

	return view
}

// RenderSingleView used to render single item's view with options
func (manager *Manager) RenderSingleView(viewName string, item models.Modellable, options *options) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	renderer := viewHandler.newHandler()
	context := makeContext(options.params)
	context.setItem(item)

	view := renderer.Render(context)

	return helpers.BuildJSON(view, options.shouldIncludeRootKey, options.rootKey)
}

// RenderListView used to render list view with options
func (manager *Manager) RenderListView(viewName string, list *models.List, options *options) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	views := make([]interface{}, list.Count())
	context := makeContext(options.params)

	for i, item := range list.Items() {
		renderer := viewHandler.newHandler()
		context.setItem(item)

		view := renderer.Render(context)
		views[i] = view
	}

	return helpers.BuildJSON(views, options.shouldIncludeRootKey, options.rootKey)
}
