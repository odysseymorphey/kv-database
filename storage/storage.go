package storage

import (
	"kv-database/models"
	"kv-database/storage/engine"
)

type Storage interface {
	Begin(query *models.Query) (*models.Result, error)
}

type storageImpl struct {
	engine engine.Engine
}

func New() Storage {
	return &storageImpl{
		engine: engine.New(),
	}
}

func (s *storageImpl) Begin(query *models.Query) (*models.Result, error) {
	res := &models.Result{}

	switch query.Command {
	case "GET":
		r, err := s.engine.Get(query.Key)
		if err != nil {
			return nil, err
		}
		res.Store(r)
	case "SET":
		s.engine.Set(query.Key, query.Value)
	case "DEL":
		err := s.engine.Delete(query.Key)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
