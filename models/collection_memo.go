package models

// FindPredicate is alias for func(Modellable) bool
type FindPredicate func(Modellable) bool

type CollectionMemo struct {
	items []Modellable
}

// makeCollection used to instantiate collection instance
func makeCollectionMemo() *CollectionMemo {
	memo := &CollectionMemo{}

	return memo
}

// Count used to count items
func (memo *CollectionMemo) Count() int {
	return len(memo.items)
}

// First used to return first item in collection
func (memo *CollectionMemo) First() Modellable {
	return memo.items[0]
}

// GetItems used to return all items in collections
func (memo *CollectionMemo) GetItems() []Modellable {
	return memo.items
}

// FindAll used to find items in collection
func (memo *CollectionMemo) FindAll(predicate FindPredicate) []Modellable {
	founds := []Modellable{}

	for _, el := range memo.items {
		if predicate(el) {
			founds = append(founds, el)
		}
	}

	return founds
}

// Find used to find item in collection, return nil if not found
func (memo *CollectionMemo) Find(predicate FindPredicate) Modellable {
	var found Modellable

	for _, el := range memo.items {
		if predicate(el) {
			found = el

			break
		}
	}

	return found
}

// FindByID used to find item in collection by id, return nil if not found
func (memo *CollectionMemo) FindByID(id string) Modellable {
	found := memo.Find(func(modellable Modellable) bool {
		return modellable.GetID() == id
	})

	return found
}

// Add used to memoize item
func (memo *CollectionMemo) Add(item Modellable) {
	memo.items = append(memo.items, item)
}
