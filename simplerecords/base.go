package simplerecords

import (
	"fmt"

	"github.com/kooinam/fabio/models"
)

// Base used to represent base classes for all models
type Base struct {
	collection          *models.Collection
	hooksHandler        *models.HooksHandler
	associationsHandler *models.AssociationsHandler
	item                models.Modellable
	ID                  string `json:"id"`
}

// InitializeBase used for setting up base attributes for a mongo record
func (base *Base) InitializeBase(context *models.Context) {
	base.collection = context.Collection()
	base.hooksHandler = context.HooksHandler()
	base.associationsHandler = context.AssociationsHandler()
	base.item = context.Item()

	base.ID = fmt.Sprintf("%v", base.collection.Adapter().(*Adapter).incrcounter())
}

// GetCollectionName used to retrieve collection's name
func (base *Base) GetCollectionName() string {
	return base.collection.Name()
}

// GetHooksHandler used to retrieve hooks handler
func (base *Base) GetHooksHandler() *models.HooksHandler {
	return base.hooksHandler
}

// GetID used to retrieve record's ID
func (base *Base) GetID() string {
	return base.ID
}

// Save used to save record in adapter
func (base *Base) Save() error {
	var err error

	err = base.hooksHandler.ExecuteValidationHooks()

	if err != nil {
		return err
	}

	return err
}

// Store used to add record to list
func (base *Base) Store() {
	base.collection.List().Add(base.item)
}

// StoreInList used to add record to selected list
func (base *Base) StoreInList(list *models.List) {
	list.Add(base.item)
}
