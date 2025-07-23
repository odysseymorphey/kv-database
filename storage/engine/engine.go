package engine

import (
	"fmt"
	"sync"
)

type Engine interface {
	Get(key string) (string, error)
	Set(key, value string)
	Delete(key string) error
}

type engineImpl struct {
	mu   sync.Mutex
	data map[string]string
}

func New() Engine {
	return &engineImpl{
		data: make(map[string]string),
	}
}

func (e *engineImpl) Get(key string) (string, error) {
	e.mu.Lock()
	v, exist := e.data[key]
	e.mu.Unlock()

	if !exist {
		return "", fmt.Errorf("key not found")
	}

	return v, nil
}

func (e *engineImpl) Set(key, value string) {
	e.mu.Lock()
	e.data[key] = value
	e.mu.Unlock()
}

func (e *engineImpl) Delete(key string) error {
	e.mu.Lock()
	if _, exist := e.data[key]; !exist {
		e.mu.Unlock()
		return fmt.Errorf("key not found")
	}

	delete(e.data, key)
	e.mu.Unlock()

	return nil
}
