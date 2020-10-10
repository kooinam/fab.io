package models

import (
	"sort"

	"github.com/kooinam/fab.io/helpers"
)

// FindPredicate is alias for func(Modellable) bool
type FindPredicate func(Modellable) bool

// SortPredictate is alias for func(int, int) bool
type SortPredictate func(int, int) bool

// List is an in-memory storage of items
type List struct {
	collection *Collection
	items      []Modellable
}

// MakeList used to instantiate list instance
func MakeList(items ...Modellable) *List {
	list := &List{
		items: []Modellable{},
	}

	for _, item := range items {
		list.Add(item)
	}

	return list
}

// MakeListWithCollection used to instantiate list instance with collection
func MakeListWithCollection(collection *Collection, items ...Modellable) *List {
	list := &List{
		collection: collection,
		items:      []Modellable{},
	}

	for _, item := range items {
		list.Add(item)
	}

	return list
}

// Count used to count items
func (list *List) Count() int {
	return len(list.items)
}

// First used to return first item in collection
func (list *List) First() Modellable {
	return list.items[0]
}

// Items used to return all items in collections
func (list *List) Items() []Modellable {
	return list.items
}

// FindAll used to find items in collection
func (list *List) FindAll(predicate FindPredicate) *List {
	newList := MakeList()
	founds := []Modellable{}

	for _, el := range list.items {
		if predicate(el) {
			founds = append(founds, el)
		}
	}

	newList.items = founds

	return newList
}

// Sort used to sort items in collection
func (list *List) Sort(predicate SortPredictate) *List {
	newList := MakeList()
	sorted := []Modellable{}

	for _, el := range list.items {
		sorted = append(sorted, el)
	}
	sort.Slice(sorted, predicate)

	newList.items = sorted

	return newList
}

// Find used to find item in collection, return nil if not found
func (list *List) Find(predicate FindPredicate) Modellable {
	var found Modellable

	for _, el := range list.items {
		if predicate(el) {
			found = el

			break
		}
	}

	return found
}

// FindByID used to find item in collection by id, return nil if not found
func (list *List) FindByID(id string) Modellable {
	found := list.Find(func(modellable Modellable) bool {
		return modellable.GetID() == id
	})

	return found
}

// Add used to store item to list
func (list *List) Add(item Modellable) {
	list.items = append(list.items, item)
}

func (list *List) Each(handler func(Modellable) bool) error {
	var err error

	items := list.Items()

	for _, item := range items {
		shouldContinue := handler(item)

		if !shouldContinue {
			break
		}
	}

	return err
}

// FirstOrCreate used to return first item in collection or create one
func (list *List) FirstOrCreate(filters helpers.H) Modellable {
	found := list.Find(func(item Modellable) bool {
		matched := true

		for key, value := range filters {
			fieldValue := helpers.GetFieldValueByName(item, key)

			if fieldValue != value {
				matched = false
			}
		}

		return matched
	})

	if found == nil {
		found = list.Create(filters)
	}

	return found
}

// Create used to create one instance
func (list *List) Create(attributes helpers.H) Modellable {
	result := list.collection.CreateWithOptions(
		attributes,
		Options().WithShouldStore(true).WithList(list),
	)

	found := result.Item()

	return found
}

func (list *List) FindIndex(id string) int {
	index := -1

	for i, item := range list.Items() {
		if item.GetID() == id {
			index = i

			break
		}
	}

	return index
}

// Destroy used to remove item from list
func (list *List) Destroy(id string) bool {
	hasDestroyed := false

	index := list.FindIndex(id)

	if index != -1 {
		list.items = append(list.items[:index], list.items[index+1:]...)
	}

	return hasDestroyed
}
