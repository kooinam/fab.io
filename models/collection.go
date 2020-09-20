package models

import (
	"fmt"

	"github.com/kooinam/fabio/helpers"
)

// NewHandler is alias for func(*Context)
type NewHandler func(*Context)

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
		list:       MakeList(),
		name:       collectionName,
		newHandler: newHandler,
	}

	return collection
}

// New used to initialize item
func (collection *Collection) New(values helpers.H) Modellable {
	attributes := helpers.MakeDictionary(values)
	context := makeContext(collection, attributes)
	collection.newHandler(context)

	if context.Item() == nil {
		panic(fmt.Sprintf("{0}'s new handler does not set item"))
	}

	item := context.Item()

	item.InitializeBase(context)
	item.GetHooksHandler().ExecuteInitializeHook(attributes)

	return item
}

// Create used to create item
func (collection *Collection) Create(values helpers.H) *SingleResult {
	result := MakeSingleResult()

	item := collection.New(values)

	err := item.Save()

	result.Set(item, err, false)

	if result.StatusSuccess() {
		item.GetHooksHandler().ExecuteAfterCreateHook()
	}

	return result
}

// CreateWithOptions used to create item with options
func (collection *Collection) CreateWithOptions(values helpers.H, options *options) *SingleResult {
	result := collection.Create(values)

	if result.StatusSuccess() {
		item := result.Item()

		if options.storable {
			if options.list == nil {
				item.Store()
			} else {
				item.StoreInList(options.list)
			}
		}
	}

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

// // WithActorizable used to configure options actorizable
// func (collection *Collection) WithActorizable(actorizable bool) *Collection {
// 	collection.options`.actorizable = actorizable

// 	return collection
// }
