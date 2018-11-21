package main

import (
	"github.com/Mushus/apartment/backend/admin"
	"github.com/Mushus/apartment/backend/service"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := service.NewDB()
	auth := service.NewAuth()

	e := echo.New()
	sg := e.Group("/session")
	sc := admin.NewSessionController(db, auth)
	sg.POST("/login", sc.Login)
	sg.POST("/logout", sc.Logout)

	gg := e.Group("/")
	gc := admin.NewGrahpQLController(db)
	gg.POST("/graphql", gc.GraphQL)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
