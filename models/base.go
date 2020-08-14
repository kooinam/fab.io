package models

// Base used to represent base classes for all models
type Base struct {
	collection *Collection
	ID         string `json:"id"`
}

func (base *Base) Initialize(collection *Collection) {
	base.collection = collection
}

func (base *Base) GetID() string {
	return base.ID
}
