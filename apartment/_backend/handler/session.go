package handler

import (
	"net/http"

	"github.com/Mushus/apartment/middleware"
	"github.com/Mushus/apartment/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Session struct {
	Auth *middleware.Auth
	DB   *gorm.DB
}

func (s Session) Login(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	var admin model.Admin
	s.DB.Where("login = ?", login).First(&admin)

	// パスワードチェック
	if admin.ComperPassword(password) {
		return c.Render(http.StatusOK, "login.tmpl", model.LoginView{})
	}

	// ログイン
	s.Auth.Login(c, admin.Login)
	return c.Redirect(http.StatusFound, "/admin")
}

func (s Session) Logout(c echo.Context) error {
	s.Auth.Logout(c)
	return c.Render(http.StatusOK, "logout.tmpl", map[string]interface{}{})
}

func (s Session) LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.tmpl", model.LoginView{})
}
