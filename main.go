package main

import (
	"kv-database/network"
	"log"
)

func main() {
	n := network.New()

	// todo: реализовать возможность конфигурации из .yaml
	// todo: сделать грейсфул шатдаун

	// todo: сделать человеческую типизацию ошибок

	// todo: добавить больше логов
	if err := n.StartListening(); err != nil {
		log.Println(err)
	}
}
