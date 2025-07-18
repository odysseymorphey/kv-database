package parser

import "strings"

type Query struct {
	Command string
	Key     string
	Value   string
}

type Parser struct {
}

func (p *Parser) Parse(query string) {
	var q Query
	splited_query := strings.Split(query, "")

	if !isCommand(splited_query[0]) {
		return // todo: сделать нормальный ретерн
	}

	q.Command = splited_query[0]

	for i := 0; i < len(splited_query); i++ {

	}
}

func isCommand(cmd string) bool {
	commands := map[string]int{
		"SET": 1,
		"GET": 1,
		"DEL": 1,
	}

	if _, ok := commands[cmd]; !ok {
		return false
	}

	return true
}
