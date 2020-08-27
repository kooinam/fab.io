package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Base used to represent base classes for all models
type Base struct {
	collection   *Collection
	hooksHandler *HooksHandler
	item         Modellable
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

// initialize used for setting up base attributes for a mongo record
func (base *Base) initialize(collection *Collection, hooksHandler *HooksHandler, item Modellable) {
	base.collection = collection
	base.hooksHandler = hooksHandler
	base.item = item
}

// GetCollectionName used to retrieve collection's name
func (base *Base) GetCollectionName() string {
	return base.collection.name
}

// getHooksHandler used to retrieve hooks handler
func (base *Base) getHooksHandler() *HooksHandler {
	return base.hooksHandler
}

// GetID used to retrieve record's ID
func (base *Base) GetID() string {
	return base.ID.String()
}

// Save used to save record in adapter
func (base *Base) Save() error {
	var err error

	err = base.hooksHandler.executeValidationHooks()

	if err != nil {
		return err
	}

	adapter := base.collection.manager.adapter

	if adapter == nil {
		err = fmt.Errorf("adapter not registered")

		return err
	}

	collection := base.collection.manager.adapter.getCollection(base.collection.name)
	ctx := adapter.getTimeoutContext()

	if base.IsNewRecord() {
		base.ID = primitive.NewObjectID()
		base.CreatedAt = time.Now()
		base.UpdatedAt = time.Now()

		results, err2 := collection.InsertOne(ctx, base.item)

		err = err2

		if err != nil {
			base.ID = primitive.NilObjectID
			base.CreatedAt = time.Time{}
			base.UpdatedAt = time.Time{}
		} else {
			base.ID = results.InsertedID.(primitive.ObjectID)
		}

	} else {
		previousUpdatedAt := base.UpdatedAt
		base.UpdatedAt = time.Now()

		results, err2 := collection.ReplaceOne(ctx, bson.M{"_id": base.ID}, base.item)

		if err2 == nil && results.MatchedCount != 1 {
			err2 = fmt.Errorf("no matched record for update")
		}

		err = err2

		if err != nil {
			base.UpdatedAt = previousUpdatedAt
		}
	}

	return err
}

// IsNewRecord used to check if record is new unsaved record
func (base *Base) IsNewRecord() bool {
	return base.ID == primitive.NilObjectID
}

// Reload used to reload record from database
func (base *Base) Reload() error {
	var err error

	// TOOD

	return err
}

// Destroy used to delete record from database
func (base *Base) Destroy() error {
	var err error

	// TOOD

	return err
}

// Memoize used to add record to memory
func (base *Base) Memoize() {
	base.collection.memo.append(base)

	base.getHooksHandler().executeAfterMemoizeHook()
}
