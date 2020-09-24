package models

// FindPredicate is alias for func(Modellable) bool
type FindPredicate func(Modellable) bool

// List is an in-memory storage of items
type List struct {
	items []Modellable
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
