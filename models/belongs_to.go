package models

type BelongsTo struct {
	collection *Collection
	key        string
	result     *SingleResult
}

func makeBelongsTo(collection *Collection) *BelongsTo {
	belongsTo := &BelongsTo{
		collection: collection,
	}

	return belongsTo
}

func (belongsTo *BelongsTo) SetKey(key string) {
	belongsTo.key = key
	belongsTo.result = belongsTo.collection.Query().Find(belongsTo.key)
}

func (belongsTo *BelongsTo) Item() Modellable {
	item := belongsTo.result.Item()

	if item == nil {
		panic("belongs_to item not found")
	}

	return item
}
