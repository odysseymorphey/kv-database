package main

import (
	"bufio"
	"kv-database/compute/parser"
	"log"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	p := parser.New()

	for {
		sc.Scan()
		if sc.Text() == "exit" {
			os.Exit(0)
		}

		// todo: сделать человеческую типизацию ошибок
		if err := p.Parse(sc.Text()); err != nil {
			log.Println(err)
		}
	}

}
