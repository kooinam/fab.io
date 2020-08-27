package models

// Manager is a singleton used to manager model's behavior
type Manager struct {
	collections []*Collection

	adapter *Adapter
}

// Setup used to setup manager
func (manager *Manager) Setup() {
	manager.collections = []*Collection{}
}

// CreateCollection used to create a collection for models manager
func (manager *Manager) CreateCollection(collectionName string, newHandler NewHandler) *Collection {
	collection := makeCollection(manager, collectionName, newHandler)

	manager.collections = append(manager.collections, collection)

	return collection
}

func (manager *Manager) RegisterAdapter(uri string, database string) error {
	var err error
	adapter, err := makeAdapter(uri, database)

	if err != nil {
		return err
	}

	manager.adapter = adapter

	return err
}
