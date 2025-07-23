package engine

import (
	"fmt"
)

type Engine interface {
	Get(key string) (string, error)
	Set(key, value string)
	Delete(key string) error
}

type engineImpl struct {
	data map[string]string
}

func New() Engine {
	return &engineImpl{
		data: make(map[string]string),
	}
}

func (e *engineImpl) Get(key string) (string, error) {
	v, ok := e.data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return v, nil
}

func (e *engineImpl) Set(key, value string) {
	e.data[key] = value
}

func (e *engineImpl) Delete(key string) error {
	if _, ok := e.data[key]; !ok {
		return fmt.Errorf("key not found")
	}

	delete(e.data, key)

	return nil
}
