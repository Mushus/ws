package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type Auth struct {
	sessionKey     string
	loginUserIDKey string
}

func NewAuth() *Auth {
	return &Auth{
		sessionKey:     "session",
		loginUserIDKey: "login_user_id",
	}
}

// 認証必須
func (m *Auth) AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(m.sessionKey, c)

		if _, ok := sess.Values[m.loginUserIDKey]; !ok {
			return c.Redirect(http.StatusFound, "/admin/login")
		}

		return next(c)
	}
}

// ログインする
func (m *Auth) Login(c echo.Context, login string) {
	sess, _ := session.Get(m.sessionKey, c)
	sess.Values[m.loginUserIDKey] = login
	sess.Save(c.Request(), c.Response())
}

// ログアウトする
func (m *Auth) Logout(c echo.Context) {
	sess, _ := session.Get(m.sessionKey, c)
	delete(sess.Values, m.loginUserIDKey)
	sess.Save(c.Request(), c.Response())
}

// ログインユーザーを取得
func (m *Auth) LoginUser(c echo.Context) (string, bool) {
	sess, _ := session.Get(m.sessionKey, c)

	loginVal, ok := sess.Values[m.loginUserIDKey]
	if !ok {
		return "", false
	}

	login, ok := loginVal.(string)
	if !ok {
		return "", false
	}

	return login, true
}
