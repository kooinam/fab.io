package fab

import (
	"github.com/kooinam/fabio/controllers"
	"github.com/kooinam/fabio/models"
	"github.com/kooinam/fabio/views"
)

type Intializer interface {
	RegisterAdapters(*models.Manager)
	RegisterCollections(*models.Manager)
	RegisterControllers(*controllers.Manager)
	RegisterViews(*views.Manager)
	Configure(*Configuration)
}
