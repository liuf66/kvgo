package main

import (
	"log"

	"github.com/liuf66/kvgo/config"
	"github.com/liuf66/kvgo/server"
)

func main() {
	server.Start(config.DefaultServerConfig)
	log.Println("kvgo started!")
}
