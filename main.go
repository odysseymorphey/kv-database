package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		sc.Scan()
		if sc.Text() == "exit" {
			os.Exit(0)
		}

		fmt.Println(sc.Text())
	}

}
