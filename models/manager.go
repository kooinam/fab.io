package models

// Manager is a singleton used to manager model's behavior
type Manager struct {
	adapters map[string]Adaptable
}

// Setup used to setup manager
func (manager *Manager) Setup() {
	manager.adapters = map[string]Adaptable{}
}

// RegisterAdapter used to register adapter
func (manager *Manager) RegisterAdapter(clientName string, adapter Adaptable) {
	manager.adapters[clientName] = adapter
}

// Adapter used to retrieve registered adapter
func (manager *Manager) Adapter(clientName string) Adaptable {
	return manager.adapters[clientName]
}

// RegisterCollection used to create a collection and register with adapter
func (manager *Manager) RegisterCollection(clientName string, collectionName string, newHandler NewHandler) *Collection {
	adapter := manager.Adapter(clientName)

	collection := makeCollection(adapter, collectionName, newHandler)

	adapter.RegisterCollection(collectionName, collection)

	return collection
}

// CreateCollection used to create a collection without registering
func (manager *Manager) CreateCollection(clientName string, collectionName string, newHandler NewHandler) *Collection {
	adapter := manager.Adapter(clientName)

	collection := makeCollection(adapter, collectionName, newHandler)

	return collection
}

// Collection used to retrieve registered collection
func (manager *Manager) Collection(clientName string, collectionName string) *Collection {
	return manager.Adapter(clientName).Collection(collectionName)
}
