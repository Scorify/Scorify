package structs

import "sync"

type SyncMap[key comparable, value any] struct {
	mutex *sync.RWMutex
	items map[key]value
}

func NewSyncMap[key comparable, value any]() *SyncMap[key, value] {
	return &SyncMap[key, value]{
		mutex: &sync.RWMutex{},
		items: make(map[key]value),
	}
}

func (m *SyncMap[key, value]) Get(k key) (value, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := m.items[k]
	return v, ok
}

func (m *SyncMap[key, value]) Set(k key, v value) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.items[k] = v
}

func (m *SyncMap[key, value]) Delete(k key) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.items, k)
}

func (m *SyncMap[key, value]) Keys() []key {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	keys := make([]key, 0, len(m.items))
	for k := range m.items {
		keys = append(keys, k)
	}

	return keys
}

func (m *SyncMap[key, value]) Map() map[key]value {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	c := make(map[key]value, len(m.items))
	for k, v := range m.items {
		c[k] = v
	}

	return c
}

func (m *SyncMap[key, value]) Legnth() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return len(m.items)
}
