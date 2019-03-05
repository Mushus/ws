package server

import (
	"net/http"

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

	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/login", loginPage)
	e.POST("/login", login)
	e.GET("/logout", logout)

	return e
}

func loginPage(c echo.Context) error {
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

func login(c echo.Context) error {
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

func logout(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
ログアウトしました。
</body>
</html>`)
}
