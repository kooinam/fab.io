package views

// Viewable is the interface for all views implementing view's functionalities
type Viewable interface {
	Render(context *Context) interface{}
}
