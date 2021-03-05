package views

import (
	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/logger"
)

type Renderer struct {
	manager              *Manager
	viewName             string
	shouldIncludeRootKey bool
	rootKey              string
	params               helpers.H
}

func makeRenderer(manager *Manager, viewName string) *Renderer {
	renderer := &Renderer{
		manager:  manager,
		viewName: viewName,
	}

	return renderer
}

func (renderer *Renderer) WithRootKey(rootKey string) *Renderer {
	renderer.shouldIncludeRootKey = true
	renderer.rootKey = rootKey

	return renderer
}

func (renderer *Renderer) WithParams(params helpers.H) *Renderer {
	renderer.params = params

	return renderer
}

// RenderSingle used to render single item's view with options
func (renderer *Renderer) RenderSingle(item interface{}) interface{} {
	viewHandler := renderer.manager.viewHandlers[renderer.viewName]

	if viewHandler == nil {
		logger.Debug("%v view not found", renderer.viewName)

		return nil
	}

	if item == nil {
		return nil
	}

	viewRenderer := viewHandler.newHandler()
	context := makeContext(renderer.manager, renderer.params)
	context.setItem(item)

	view := viewRenderer.Render(context)

	return helpers.BuildJSON(view, renderer.shouldIncludeRootKey, renderer.rootKey)
}

// RenderList used to render list view with options
func (renderer *Renderer) RenderList(items []interface{}) interface{} {
	viewHandler := renderer.manager.viewHandlers[renderer.viewName]

	if viewHandler == nil {
		logger.Debug("%v view not found", renderer.viewName)

		return nil
	}

	views := make([]interface{}, len(items))
	context := makeContext(renderer.manager, renderer.params)

	for i, item := range items {
		viewRenderer := viewHandler.newHandler()
		context.setItem(item)

		view := viewRenderer.Render(context)
		views[i] = view
	}

	return helpers.BuildJSON(views, renderer.shouldIncludeRootKey, renderer.rootKey)
}
