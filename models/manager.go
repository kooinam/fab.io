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

// RegisterCollection used to create a collection for models manager
func (manager *Manager) RegisterCollection(clientName string, collectionName string, newHandler NewHandler) *Collection {
	adapter := manager.Adapter(clientName)

	collection := makeCollection(adapter, collectionName, newHandler)

	adapter.RegisterCollection(collection)

	return collection
}
