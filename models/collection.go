package models

import (
	"fmt"

	"github.com/kooinam/fabio/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

// NewHandler is alias for func(args ...interface{}) Modellable
type NewHandler func(*Collection, *HooksHandler) Modellable

// Collection used to contain models
type Collection struct {
	manager    *Manager
	cacher     *CollectionCacher
	name       string
	newHandler NewHandler
}

// makeCollection used to instantiate collection instance
func makeCollection(manager *Manager, collectionName string, newHandler NewHandler) *Collection {
	collection := &Collection{
		manager:    manager,
		cacher:     makeCollectionCacher(),
		name:       collectionName,
		newHandler: newHandler,
	}

	return collection
}

// New used to initialize item
func (collection *Collection) New(values helpers.H) Modellable {
	hooksHandler := makeHooksHandler()
	item := collection.newHandler(collection, hooksHandler)

	item.Instantiate(collection, hooksHandler, item)

	item.GetHooksHandler().executeInitializeHook(helpers.MakeDictionary(values))

	return item
}

// Create used to create item
func (collection *Collection) Create(values helpers.H) (Modellable, error) {
	var err error

	item := collection.New(values)

	err = item.Save()

	return item, err
}

// Query returns query wrapper for retrieving records in adapter
func (collection *Collection) Query() *Query {
	var err error
	var query *Query
	adapter := collection.manager.adapter

	if adapter == nil {
		err = fmt.Errorf("adapter not registered")

		query = &Query{
			err: err,
		}
	} else {
		query = &Query{
			adapter:    adapter,
			collection: collection,
			filters:    bson.M{},
		}
	}

	return query
}
