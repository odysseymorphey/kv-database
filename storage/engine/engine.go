package engine

import (
	"fmt"
	"kv-database/models"
)

type Engine interface {
	Begin(query models.Query) (string, error)
}

type engineImpl struct {
	data map[string]string
}

func New() Engine {
	return &engineImpl{
		data: make(map[string]string),
	}
}

func (e *engineImpl) Begin(query models.Query) (string, error) {
	switch query.Command {
	case "GET":
		return e.get(query.Key)
	case "SET":
		e.set(query.Key, query.Value)
	case "DEL":
		err := e.delete(query.Key)
		return "", err
	}

	return "", nil
}

func (e *engineImpl) get(key string) (string, error) {
	v, ok := e.data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return v, nil
}

func (e *engineImpl) set(key, value string) {
	e.data[key] = value
}

func (e *engineImpl) delete(key string) error {
	if _, ok := e.data[key]; !ok {
		return fmt.Errorf("key not found")
	}

	delete(e.data, key)

	return nil
}
