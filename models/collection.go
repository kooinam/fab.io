package models

// CreateHandler is alias for func(args ...interface{}) Modellable
type CreateHandler func(collection *Collection, args ...interface{}) Modellable

// FindPredicate is alias for func(Modellable) bool
type FindPredicate func(Modellable) bool

// Collection used to contain models
type Collection struct {
	name          string
	createHandler CreateHandler
	items         []Modellable
}

// makeCollection used to instantiate collection instance
func makeCollection(collectionName string, createHandler CreateHandler) *Collection {
	collection := &Collection{
		name:          collectionName,
		createHandler: createHandler,
	}

	return collection
}

// Count used to count items
func (collection *Collection) Count() int {
	return len(collection.items)
}

// Create used to create item
func (collection *Collection) Create(args ...interface{}) Modellable {
	item := collection.createHandler(collection, args...)

	return item
}

// Find used to find item in collection, return nil if not found
func (collection *Collection) Find(predicate FindPredicate) Modellable {
	var found Modellable

	for _, el := range collection.items {
		if predicate(el) {
			found = el

			break
		}
	}

	return found
}

// FindByID used to find item in collection by id, return nil if not found
func (collection *Collection) FindByID(id string) Modellable {
	found := collection.Find(func(modellable Modellable) bool {
		return modellable.GetID() == id
	})

	return found
}

// FindOrCreate used to find item in collection, create one if not found
func (collection *Collection) FindOrCreate(predicate FindPredicate) Modellable {
	found := collection.Find(predicate)

	if found == nil {
		found = collection.Create()
	}

	return found
}

// First used to return first item in collection
func (collection *Collection) First() Modellable {
	return collection.items[0]
}

// GetItems used to return all items in collections
func (collection *Collection) GetItems() []Modellable {
	return collection.items
}

func (collection *Collection) append(item Modellable) {
	collection.items = append(collection.items, item)
}
