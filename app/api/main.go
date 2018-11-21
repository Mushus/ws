package main

import (
	"github.com/Mushus/app/api/config"
	"github.com/Mushus/app/api/server"
)

func main() {
	cfg := config.LoadConfig()

	s := server.NewServer(cfg)
	s.Start()
}
