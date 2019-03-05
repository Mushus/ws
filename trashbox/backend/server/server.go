package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server サーバーインスタンスです
type Server struct {
	router *echo.Echo
	db     *DB
}

// New サーバーを作成する
func New() (*Server, error) {

	r := createRouter()

	db, err := NewDB("trashbox.db")
	if err != nil {
		r.Logger.Fatal(err)
		return nil, err
	}

	return &Server{
		router: r,
		db:     db,
	}, nil
}

// Start サーバーを起動します
func (s Server) Start() {
	addr := ":8080"

	r := s.router
	if err := r.Start(addr); err != nil {
		r.Logger.Fatal(err)
	}
}

func createRouter() *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Validator = NewValidator()

	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/login", loginPage)
	e.POST("/login", login)
	e.GET("/logout", logout)

	return e
}
