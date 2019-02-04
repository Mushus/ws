package main // import "github.com/Mushus/gate"

import (
	"log"

	"github.com/Mushus/gate/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.Load("./gate.yml")
	if err != nil {
		log.Fatalln(err)
	}
	port := cfg.App.Port.HTTPString()
	services := cfg.Services

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	for _, service := range services {
		path := service.Key
		options := service.Value
		log.Printf("%#v", path)
		log.Printf("%#v", options)
	}

	err = e.Start(port)
	if err != nil {
		log.Fatalln(err)
	}
}
