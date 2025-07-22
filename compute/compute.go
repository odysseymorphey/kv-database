package compute

import "kv-database/compute/parser"

type Compute interface {
	Parse(query string) (string, error)
}

type computeImpl struct {
	parser parser.Parser
}

func New() Compute {
	return &computeImpl{
		parser: parser.New(),
	}
}

func (c *computeImpl) Parse(query string) (string, error) {
	return c.parser.Parse(query)
}
