package models

// FindPredicate is alias for func(Modellable) bool
type FindPredicate func(Modellable) bool

type CollectionCacher struct {
	items []Modellable
}

// makeCollection used to instantiate collection instance
func makeCollectionCacher() *CollectionCacher {
	cacher := &CollectionCacher{}

	return cacher
}

// Count used to count items
func (cacher *CollectionCacher) Count() int {
	return len(cacher.items)
}

// First used to return first item in collection
func (cacher *CollectionCacher) First() Modellable {
	return cacher.items[0]
}

// GetItems used to return all items in collections
func (cacher *CollectionCacher) GetItems() []Modellable {
	return cacher.items
}

// FindAll used to find items in collection
func (cacher *CollectionCacher) FindAll(predicate FindPredicate) []Modellable {
	founds := []Modellable{}

	for _, el := range cacher.items {
		if predicate(el) {
			founds = append(founds, el)
		}
	}

	return founds
}

// Find used to find item in collection, return nil if not found
func (cacher *CollectionCacher) Find(predicate FindPredicate) Modellable {
	var found Modellable

	for _, el := range cacher.items {
		if predicate(el) {
			found = el

			break
		}
	}

	return found
}

// FindByID used to find item in collection by id, return nil if not found
func (cacher *CollectionCacher) FindByID(id string) Modellable {
	found := cacher.Find(func(modellable Modellable) bool {
		return modellable.GetID() == id
	})

	return found
}

func (cacher *CollectionCacher) append(item Modellable) {
	cacher.items = append(cacher.items, item)
}
