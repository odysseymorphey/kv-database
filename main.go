package main

import (
	"bufio"
	"fmt"
	"kv-database/compute"
	"log"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	cmp := compute.New()

	for {
		sc.Scan()
		if sc.Text() == "exit" {
			os.Exit(0)
		}

		// todo: сделать человеческую типизацию ошибок

		// todo: сделать нормальный принтер
		r, err := cmp.Exec(sc.Text())
		if err != nil {
			log.Println(err)
		} else {
			if r.String() != "" {
				fmt.Println(r)
			}
		}
	}

}
