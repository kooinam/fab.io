package models

// Collection used to contain models
type Collection struct {
	createHandler func(args ...interface{}) Base
	items         []Base
}

// MakeCollection used to instantiate collection instance
func MakeCollection(createHandler func(args ...interface{}) Base) *Collection {
	collection := &Collection{
		createHandler: createHandler,
	}

	return collection
}

// Count used to count items
func (collection *Collection) Count() int {
	return len(collection.items)
}

// Create used to create item
func (collection *Collection) Create(args ...interface{}) Base {
	item := collection.createHandler(args...)

	collection.append(item)

	return item
}

// Find used to find item in collection, return nil if not found
func (collection *Collection) Find(predicate func(Base) bool) Base {
	var found Base

	for _, el := range collection.items {
		if predicate(el) {
			found = el

			break
		}
	}

	return found
}

// FindByID used to find item in collection by id, return nil if not found
func (collection *Collection) FindByID(id string) Base {
	found := collection.Find(func(base Base) bool {
		return base.GetID() == id
	})

	return found
}

// FindOrCreate used to find item in collection, create one if not found
func (collection *Collection) FindOrCreate(predicate func(Base) bool) Base {
	found := collection.Find(predicate)

	if found == nil {
		found = collection.Create()
	}

	return found
}

// First used to return first item in collection
func (collection *Collection) First() Base {
	return collection.items[0]
}

// GetItems used to return all items in collections
func (collection *Collection) GetItems() []Base {
	return collection.items
}

func (collection *Collection) append(item Base) {
	collection.items = append(collection.items, item)
}
