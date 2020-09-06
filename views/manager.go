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
func (manager *Manager) RenderView(viewName string, params helpers.H, rootKey string) interface{} {
	view := manager.RenderViewWithoutRootKey(viewName, params)

	return helpers.BuildJSON(view, true, rootKey)
}

// RenderViewWithoutRootKey used to render view without root key
func (manager *Manager) RenderViewWithoutRootKey(viewName string, params helpers.H) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	renderer := viewHandler.newHandler()
	context := makeContext(params)

	view := renderer.Render(context)

	return view
}

// RenderSingleView used to render single item's view
func (manager *Manager) RenderSingleView(viewName string, item models.Modellable, params helpers.H, rootKey string) interface{} {
	view := manager.RenderSingleViewWithoutRootKey(viewName, item, params)

	return helpers.BuildJSON(view, true, rootKey)
}

// RenderSingleViewWithoutRootKey used to render single item's view without root key
func (manager *Manager) RenderSingleViewWithoutRootKey(viewName string, item models.Modellable, params helpers.H) interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	renderer := viewHandler.newHandler()
	context := makeContext(params)
	context.setItem(item)

	view := renderer.Render(context)

	return view
}

// RenderListView used to render list view
func (manager *Manager) RenderListView(viewName string, list *models.List, params helpers.H, rootKey string) interface{} {
	view := manager.RenderListViewWithoutRootKey(viewName, list, params)

	return helpers.BuildJSON(view, true, rootKey)
}

// RenderListViewWithoutRootKey used to render list view without root key
func (manager *Manager) RenderListViewWithoutRootKey(viewName string, list *models.List, params helpers.H) []interface{} {
	viewHandler := manager.viewHandlers[viewName]

	if viewHandler == nil {
		return nil
	}

	views := make([]interface{}, list.Count())
	context := makeContext(params)
	for i, item := range list.Items() {
		renderer := viewHandler.newHandler()
		context.setItem(item)

		view := renderer.Render(context)
		views[i] = view
	}

	return views
}
