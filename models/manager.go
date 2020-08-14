package models

type Manager struct {
}

func (manager *Manager) Setup() {

}

func (manager *Manager) CreateCollection(collectionName string, createHandler CreateHandler) *Collection {
	collection := makeCollection(collectionName, createHandler)

	return collection
}
