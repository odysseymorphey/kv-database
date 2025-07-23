package parser

import (
	"fmt"
	"kv-database/models"
	"kv-database/storage"
	"strings"
)

type Parser interface {
	Parse(query string) (*models.Query, error)
}

type parserImpl struct {
	storage storage.Storage
}

func New() Parser {
	return &parserImpl{
		storage: storage.New(),
	}
}

func (p *parserImpl) Parse(query string) (*models.Query, error) {
	if query == "" {
		return nil, fmt.Errorf("empty querry string")
	}

	q := &models.Query{}
	splitQuery := strings.Fields(query)

	if !isCommand(splitQuery[0]) {
		return nil, fmt.Errorf("wrong command. available commands: GET, SET, DEL")
	}

	q.Command = splitQuery[0]

	switch len(splitQuery) {
	case 3:
		q.Key = splitQuery[1]
		q.Value = splitQuery[2]
	case 2:
		q.Key = splitQuery[1]
	default:
		return nil, fmt.Errorf("invalid query string")
	}

	return q, nil
}

func isCommand(cmd string) bool {
	commands := map[string]struct{}{
		"SET": {},
		"GET": {},
		"DEL": {},
	}

	if _, ok := commands[cmd]; !ok {
		return false
	}

	return true
}
