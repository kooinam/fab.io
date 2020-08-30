package models

import (
	"github.com/kooinam/fabio/helpers"
)

// NewHandler is alias for func(args ...interface{}) Modellable
type NewHandler func(*Collection, *HooksHandler) Modellable

// Collection used to contain models
type Collection struct {
	adapter    Adaptable
	list       *List
	name       string
	newHandler NewHandler
}

// makeCollection used to instantiate collection instance
func makeCollection(adapter Adaptable, collectionName string, newHandler NewHandler) *Collection {
	collection := &Collection{
		adapter:    adapter,
		list:       makeList(),
		name:       collectionName,
		newHandler: newHandler,
	}

	return collection
}

// New used to initialize item
func (collection *Collection) New(values helpers.H) Modellable {
	hooksHandler := makeHooksHandler()
	item := collection.newHandler(collection, hooksHandler)

	item.InitializeBase(collection, hooksHandler, item)

	item.GetHooksHandler().ExecuteInitializeHook(helpers.MakeDictionary(values))

	return item
}

// Create used to create item
func (collection *Collection) Create(values helpers.H) *SingleResult {
	result := MakeSingleResult()

	item := collection.New(values)

	err := item.Save()

	result.Set(item, err, false)

	return result
}

// Query returns query wrapper for retrieving records in adapter
func (collection *Collection) Query() Queryable {
	adapter := collection.Adapter()

	query := adapter.NewQuery(collection)

	return query
}

// Adapter used to retrieve collection's adapter
func (collection *Collection) Adapter() Adaptable {
	adapter := collection.adapter

	if adapter == nil {
		panic("adapter not registered")
	}

	return adapter
}

// Name used to retrieve collection's name
func (collection *Collection) Name() string {
	return collection.name
}

// List used to retrieve in-memory list
func (collection *Collection) List() *List {
	return collection.list
}
