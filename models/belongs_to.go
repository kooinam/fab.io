package models

import "github.com/kooinam/fab.io/helpers"

type BelongsTo struct {
	collection *Collection
	key        string
	result     *SingleResult
	foreignKey string
}

func makeBelongsTo(collection *Collection) *BelongsTo {
	belongsTo := &BelongsTo{
		collection: collection,
		foreignKey: "ID",
	}

	return belongsTo
}

func (belongsTo *BelongsTo) WithForeignKey(foreignKey string) *BelongsTo {
	belongsTo.foreignKey = foreignKey

	return belongsTo
}

func (belongsTo *BelongsTo) SetKey(key string) {
	belongsTo.key = key

	if belongsTo.foreignKey == "ID" {
		belongsTo.result = belongsTo.collection.Query().Find(belongsTo.key)
	} else {
		filters := helpers.H{}
		filters[belongsTo.foreignKey] = belongsTo.key

		belongsTo.result = belongsTo.collection.Query().Where(filters).First()
	}
}

func (belongsTo *BelongsTo) Key() string {
	return belongsTo.key
}

func (belongsTo *BelongsTo) Item() Modellable {
	if belongsTo.result == nil {
		return nil
	}

	item := belongsTo.result.Item()

	return item
}
