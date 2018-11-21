package main

import (
	"log"

	"github.com/Mushus/app/facade/config"
	"github.com/Mushus/app/facade/server"
)

func main() {
	cfg := config.LoadConfig()

	s, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("failed to set up server: %v", err)
	}
	s.Start()
}
