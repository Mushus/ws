package main

import (
	"log"

	"github.com/Mushus/trashbox/backend/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	s.Start()
}
