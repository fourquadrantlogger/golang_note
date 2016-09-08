package main

import (
	"sync"
)

type SafeMap struct {
	lock *sync.RWMutex
	sm   map[interface{}]interface{}
}

func SafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		sm:   make(map[interface{}]interface{}),
	}
}

//Get from maps return the k's value
func (m *SafeMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.sm[k]; ok {
		return val
	}
	return nil
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *SafeMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.sm[k]; !ok {
		m.sm[k] = v
	} else if val != v {
		m.sm[k] = v
	} else {
		return false
	}
	return true
}

// Returns true if k is exist in the map.
func (m *SafeMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.sm[k]; !ok {
		return false
	}
	return true
}

func (m *SafeMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.sm, k)
}
