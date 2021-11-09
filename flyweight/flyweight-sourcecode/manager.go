package main

//Manager struct
type Manager struct {
	id   string
	name string
	dept string
}

func NewManager(id string) *Manager {
	return &Manager{id: id}
}

//Manager class method getId
func (manager Manager) getId() string {
	return manager.id
}
