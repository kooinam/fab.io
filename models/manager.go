package models

type Manager struct {
}

func (manager *Manager) Setup() {

}

func (manager *Manager) RegisterCollection(collectionName string, createHandler CreateHandler) *Collection {
	collection := makeCollection(collectionName, createHandler)

	return collection
}
