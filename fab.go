package fab

import (
	"github.com/kooinam/fabio/actors"
	"github.com/kooinam/fabio/controllers"
	"github.com/kooinam/fabio/models"
)

var engine *Engine

// Engine is the core engine for fab.io
type Engine struct {
	controllerManager *controllers.Manager
	actorManager      *actors.Manager
	modelManager      *models.Manager
}

// Setup used to setup engine
func Setup() {
	engine = &Engine{
		controllerManager: &controllers.Manager{},
		actorManager:      &actors.Manager{},
		modelManager:      &models.Manager{},
	}

	engine.controllerManager.Setup()
	engine.actorManager.Setup()

	serveStats()
}

// ControllerManager used to retrieve controller manager
func ControllerManager() *controllers.Manager {
	return engine.controllerManager
}

// ActorManager used to retrieve actor manager
func ActorManager() *actors.Manager {
	return engine.actorManager
}

// ModelManager used to retrieve model manager
func ModelManager() *models.Manager {
	return engine.modelManager
}
