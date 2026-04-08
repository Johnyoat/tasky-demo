package main

import (
	"log"

	"github.com/johnyoat/tasky-demo/tasky-api/internal/api"
)

func main() {

	server := api.NewServer()
	log.Fatal(server.Start())
}
