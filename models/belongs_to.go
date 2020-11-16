package models

import (
	"fmt"

	"github.com/kooinam/fab.io/helpers"
	"github.com/kooinam/fab.io/logger"
)

type BelongsTo struct {
	collection *Collection
	item       Modellable
	key        string
	result     *SingleResult
	foreignKey string
	primaryKey string
}

func makeBelongsTo(collection *Collection) *BelongsTo {
	belongsTo := &BelongsTo{
		collection: collection,
		primaryKey: "ID",
	}

	return belongsTo
}

// WithPrimaryKey used to set belongs_to primary key
func (belongsTo *BelongsTo) WithPrimaryKey(primaryKey string) *BelongsTo {
	belongsTo.primaryKey = primaryKey

	return belongsTo
}

// WithForeignKey used to set belongs_to foreign key
func (belongsTo *BelongsTo) WithForeignKey(foreignKey string) *BelongsTo {
	belongsTo.foreignKey = foreignKey

	return belongsTo
}

// Clear used to clear association
func (belongsTo *BelongsTo) Clear() error {
	var err error

	belongsTo.key = ""
	if belongsTo.item != nil && len(belongsTo.foreignKey) > 0 {
		helpers.SetFieldValueByNameStr(belongsTo.item, belongsTo.foreignKey, "")
	}

	belongsTo.result = nil

	return err
}

// Set used to set association
func (belongsTo *BelongsTo) Set(item Modellable) error {
	var err error

	key := helpers.GetFieldValueByName(item, belongsTo.foreignKey)
	belongsTo.key = key.(string)

	if belongsTo.item != nil && len(belongsTo.foreignKey) > 0 {
		helpers.SetFieldValueByNameStr(belongsTo.item, belongsTo.foreignKey, belongsTo.key)
	}

	result := MakeSingleResult()
	result.Set(item, nil, false)

	belongsTo.result = result

	return err
}

// SetKey used to set association's reference key
func (belongsTo *BelongsTo) SetKey(key string) error {
	var err error

	belongsTo.key = key
	if belongsTo.item != nil && len(belongsTo.foreignKey) > 0 {
		helpers.SetFieldValueByNameStr(belongsTo.item, belongsTo.foreignKey, belongsTo.key)
	}

	if belongsTo.primaryKey == "ID" {
		belongsTo.result = belongsTo.collection.Query().Find(belongsTo.key)
	} else {
		filters := helpers.H{}
		filters[belongsTo.primaryKey] = belongsTo.key

		belongsTo.result = belongsTo.collection.Query().Where(filters).First()
	}

	err = belongsTo.result.Error()
	if err != nil {
		return err
	}

	if belongsTo.IsEmpty() {
		err = fmt.Errorf("belongs_to not found - %v:%v:%v", belongsTo.collection.Name(), belongsTo.key, belongsTo.result.Error())

		return err
	}

	return err
}

// Key used to retrieve association's reference key
func (belongsTo *BelongsTo) Key() string {
	return belongsTo.key
}

// Item used to retrieve association's item
func (belongsTo *BelongsTo) Item() Modellable {
	if belongsTo.result == nil {
		return nil
	}

	if belongsTo.result.Item() == nil {
		logger.Debug("belongs_to not found - %v:%v:%v", belongsTo.collection.Name(), belongsTo.key, belongsTo.result.Error())

		return nil
	}

	item := belongsTo.result.Item()

	return item
}

// IsEmpty used to determine if belongs_to association is empty
func (belongsTo *BelongsTo) IsEmpty() bool {
	if belongsTo.result == nil || belongsTo.result.Item() == nil {
		return true
	}

	return false
}

// Equals used to compare if item is association
func (belongsTo *BelongsTo) Equals(item Modellable) bool {
	return item.GetID() == belongsTo.Key()
}
