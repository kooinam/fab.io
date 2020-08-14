package models

import (
	"fmt"
)

// Base used to represent base classes for all models
type Base struct {
	collection *Collection
	ID         string `json:"id"`
}

func (base *Base) Initialize(collection *Collection) {
	base.collection = collection
	base.ID = fmt.Sprintf("%v", collection.Count()+1)
}

func (base *Base) GetCollectionName() string {
	return base.collection.name
}

func (base *Base) GetID() string {
	return base.ID
}
