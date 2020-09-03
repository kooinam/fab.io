package views

import (
	"github.com/kooinam/fabio/helpers"
)

// Viewable is the interface for all views implementing view's functionalities
type Viewable interface {
	Render(*helpers.Dictionary) interface{}
}
