package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jmoiron/sqlx"

	"golang.org/x/xerrors"
)

// Server サーバーインスタンスです
type Server struct {
	router *echo.Echo
	db *sqlx.DB
}

// New サーバーを作成する
func New() (*Server, error) {

	r := createRouter()
	db, err := createDatabase()
	if err != nil {
		r.Logger.Fatal(err)
		return nil, err
	}

	return &Server{
		router: r,
		db: db,
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
	e.Use(middleware.Logger(), middleware.Recover())
	e.GET("/login", getLogin)
	return e
}

func createDatabase() (*sqlx.DB, error) {
	workspace, err := os.Getwd()
	if err != nil {
		return nil, xerrors.Errorf("failed to get current directory: %w", err)
	}
	// データベースの保存先
	dbPath := filepath.Join(workspace, "trashbox.db")

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, xerrors.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

func getLogin(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
<body>
<form method="POST" action="login">
<input type="text" name="user" placeholder="user name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
</body>
</html>`)
}

func postLogin(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
<body>
<form method="POST" action="login">
<input type="text" name="user" placeholder="user name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
</body>
</html>`)
}
