package models

type AssociationsHandler struct {
	belongsTos []*BelongsTo
}

func makeAssociationsHandler() *AssociationsHandler {
	handler := &AssociationsHandler{}

	return handler
}

func (handler *AssociationsHandler) RegisterBelongsTo(collection *Collection) *BelongsTo {
	belongsTo := makeBelongsTo(collection)

	return belongsTo
}
