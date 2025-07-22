package storage

import (
	"kv-database/models"
	"kv-database/storage/engine"
)

type Storage interface {
	Begin(query models.Query) (string, error)
}

type storageImpl struct {
	engine engine.Engine
}

func New() Storage {
	return &storageImpl{
		engine: engine.New(),
	}
}

func (s *storageImpl) Begin(query models.Query) (string, error) {
	return s.engine.Begin(query)
}
