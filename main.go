package main

import (
	"github.com/santoshkc89/inventory_management/webserver"
)

func main() {

	server := webserver.Server{
		Address: "localhost",
		Port:    8080,
	}

	server.Run()
}
