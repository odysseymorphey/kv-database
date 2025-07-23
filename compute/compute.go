package compute

import (
	"kv-database/compute/parser"
	"kv-database/models"
	"kv-database/storage"
)

type Compute interface {
	Exec(query string) (*models.Result, error)
}

type computeImpl struct {
	parser  parser.Parser
	storage storage.Storage
}

func New() Compute {
	return &computeImpl{
		parser:  parser.New(),
		storage: storage.New(),
	}
}

func (c *computeImpl) Exec(queryStr string) (*models.Result, error) {
	q, err := c.parser.Parse(queryStr)
	if err != nil {
		return nil, err
	}

	r, err := c.storage.Begin(q)
	if err != nil {
		return nil, err
	}
	return r, nil
}
