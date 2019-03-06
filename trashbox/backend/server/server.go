package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/xerrors"
)

// Server サーバーインスタンスです
type Server struct {
	router *echo.Echo
	db     *DB
}

// New サーバーを作成する
func New() (*Server, error) {

	db, err := NewDB("trashbox.db")
	if err != nil {
		return nil, err
	}

	handler := NewHandler(db)

	r, err := createRouter(handler)
	if err != nil {
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
	// let's start
	if err := r.Start(addr); err != nil {
		r.Logger.Fatal(err)
	}
}

func createRouter(handler Handler) (*echo.Echo, error) {
	e := echo.New()

	// settings
	e.HideBanner = true
	e.HidePort = true
	e.Validator = NewValidator()
	e.Renderer = NewRenderer()

	// set up middleware
	session, err := NewSession()
	if err != nil {
		return nil, xerrors.Errorf("failed to setup session: %w", err)
	}
	logger := middleware.Logger()
	recover := middleware.Recover()
	e.Use(session, logger, recover)

	// set up routings
	e.GET("/login", handler.GetLogin)
	e.POST("/login", handler.PostLogin)
	e.GET("/logout", handler.GetLogout)

	return e, nil
}
