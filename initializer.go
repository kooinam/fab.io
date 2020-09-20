package fab

import (
	"github.com/kooinam/fab.io/controllers"
	"github.com/kooinam/fab.io/models"
	"github.com/kooinam/fab.io/views"
)

type Intializer interface {
	// Configure used to configure configurations like port and httphandler
	Configure(*Configuration)

	// RegisterAdapters used to register adapters
	RegisterAdapters(*models.Manager)
	// RegisterCollections used to register collections
	RegisterCollections(*models.Manager)
	// RegisterControllers used to register controllers
	RegisterControllers(*controllers.Manager)
	// RegisterViews used to register views
	RegisterViews(*views.Manager)

	//BeforeServe used for custom initializations after fab.io initializes and before serving like loading some seed application data
	BeforeServe()
}
