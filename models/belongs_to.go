package models

import (
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

func (belongsTo *BelongsTo) ClearKey() error {
	var err error

	belongsTo.key = ""
	if belongsTo.item != nil && len(belongsTo.foreignKey) > 0 {
		helpers.SetFieldValueByNameStr(belongsTo.item, belongsTo.foreignKey, "")
	}

	belongsTo.result = nil

	return err
}

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

	return err
}

func (belongsTo *BelongsTo) Key() string {
	return belongsTo.key
}

func (belongsTo *BelongsTo) Item() Modellable {
	if belongsTo.result == nil {
		logger.Debug("belongs_to not found - %v:%v", belongsTo.collection.Name(), belongsTo.key)

		return nil
	}

	item := belongsTo.result.Item()

	return item
}
