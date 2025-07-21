package engine

import (
	"fmt"
	"kv-database/models"
)

type Engine struct {
	data map[string]string
}

func New() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}

func (e *Engine) Begin(query models.Query) (string, error) {
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

func (e *Engine) get(key string) (string, error) {
	v, ok := e.data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return v, nil
}

func (e *Engine) set(key, value string) {
	e.data[key] = value
}

func (e *Engine) delete(key string) error {
	if _, ok := e.data[key]; !ok {
		return fmt.Errorf("key not found")
	}

	delete(e.data, key)

	return nil
}
