package main

import (
	"github.com/Mushus/trashbox/backend/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s, err := server.New()
	if err != nil {
		return
	}
	s.Start()
}
