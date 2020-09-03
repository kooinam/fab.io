package views

type ViewHandler struct {
	newHandler func() Viewable
}

func makeViewHandler(newHandler func() Viewable) *ViewHandler {
	viewHandler := &ViewHandler{
		newHandler: newHandler,
	}

	return viewHandler
}
