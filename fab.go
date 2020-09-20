package fab

import (
	"github.com/kooinam/fab.io/actors"
	"github.com/kooinam/fab.io/controllers"
	"github.com/kooinam/fab.io/models"
	"github.com/kooinam/fab.io/views"
)

var engine *Engine

// Engine is the core engine for fab.io
type Engine struct {
	modelManager      *models.Manager
	controllerManager *controllers.Manager
	viewManager       *views.Manager
	actorManager      *actors.Manager
	configuration     *Configuration
}

// Setup used to setup engine
func Setup() {
	engine = &Engine{
		modelManager:      &models.Manager{},
		controllerManager: &controllers.Manager{},
		viewManager:       &views.Manager{},
		actorManager:      &actors.Manager{},
	}

	engine.modelManager.Setup()
	engine.viewManager.Setup()
	engine.controllerManager.Setup(ViewManager())
	engine.actorManager.Setup()

	serveStats()
}

// ConfigureAndServe used to setup applications and start server
// register adapters, collections for modelmanager
// register controllers, routings for controllermanager
// register views for viewmanager
// configure configuration
func ConfigureAndServe(initializer Intializer) {
	Setup()

	engine.configuration = &Configuration{}
	initializer.Configure(engine.configuration)

	initializer.RegisterAdapters(ModelManager())
	initializer.RegisterCollections(ModelManager())
	initializer.RegisterControllers(ControllerManager())
	initializer.RegisterViews(ViewManager())

	initializer.BeforeServe()

	ControllerManager().Serve(engine.configuration.port, engine.configuration.httpHandler)
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

// ViewManager used to retrieve view manager
func ViewManager() *views.Manager {
	return engine.viewManager
}
