package main

import (
	"github.com/Mushus/app/stdapp/server/http/appsettings"
	"github.com/Mushus/app/stdapp/server/http/usersettings"
	"github.com/Mushus/app/stdapp/server/middleware"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	s := e.Group("/.server")
	s.GET("/.server/app-settings", appsettings.Handler)
	s.GET("/.server/healthcheck", healthcheck.NewHanlder())
	s.GET("/.user-settings/", middleware.HandlerWithSession(usersettings.NewHandler()))
	e.Logger.Fatal(e.Start(":8080"))
}
