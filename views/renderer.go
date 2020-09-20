package views

import (
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/logger"
	"github.com/kooinam/fabio/models"
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
func (renderer *Renderer) RenderSingle(item models.Modellable) interface{} {
	viewHandler := renderer.manager.viewHandlers[renderer.viewName]

	if viewHandler == nil {
		logger.Debug("%v view not found", renderer.viewName)

		return nil
	}

	viewRenderer := viewHandler.newHandler()
	context := makeContext(renderer.manager, renderer.params)
	context.setItem(item)

	view := viewRenderer.Render(context)

	return helpers.BuildJSON(view, renderer.shouldIncludeRootKey, renderer.rootKey)
}

// RenderList used to render list view with options
func (renderer *Renderer) RenderList(list *models.List) interface{} {
	viewHandler := renderer.manager.viewHandlers[renderer.viewName]

	if viewHandler == nil {
		logger.Debug("%v view not found", renderer.viewName)

		return nil
	}

	views := make([]interface{}, list.Count())
	context := makeContext(renderer.manager, renderer.params)

	for i, item := range list.Items() {
		viewRenderer := viewHandler.newHandler()
		context.setItem(item)

		view := viewRenderer.Render(context)
		views[i] = view
	}

	return helpers.BuildJSON(views, renderer.shouldIncludeRootKey, renderer.rootKey)
}
