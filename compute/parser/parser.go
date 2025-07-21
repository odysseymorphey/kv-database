package parser

import (
	"fmt"
	"kv-database/models"
	"kv-database/storage/engine"
	"strings"
)

type Parser struct {
	engine *engine.Engine
}

func New() *Parser {
	return &Parser{
		engine: engine.New(),
	}
}

func (p *Parser) Parse(query string) error {
	if query == "" {
		return fmt.Errorf("empty querry string")
	}

	var q models.Query
	splitQuery := strings.Split(query, "")

	if !isCommand(splitQuery[0]) {
		return fmt.Errorf("wrong command. available commands: GET, SET, DEL")
	}

	q.Command = splitQuery[0]

	switch len(splitQuery) {
	case 3:
		q.Key = splitQuery[1]
		q.Value = splitQuery[2]
	case 2:
		q.Key = splitQuery[1]
	default:
		return fmt.Errorf("invalid query string")
	}

	return nil
}

func isCommand(cmd string) bool {
	commands := map[string]interface{}{
		"SET": struct{}{},
		"GET": struct{}{},
		"DEL": struct{}{},
	}

	if _, ok := commands[cmd]; !ok {
		return false
	}

	return true
}
