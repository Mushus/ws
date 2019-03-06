package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Handler handler
type Handler struct {
	db *DB
}

// NewHandler ハンドラを生成する
func NewHandler(db *DB) Handler {
	return Handler{
		db: db,
	}
}

// GetLogin ログインページ
func (h Handler) GetLogin(c echo.Context) error {
	return c.Render(http.StatusOK, TmplLogin, LoginView{
		Errors: ValidationResult{},
	})
}

// LoginParam ログイン
type LoginParam struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

// PostLogin ログイン処理
func (h Handler) PostLogin(c echo.Context) error {
	var prm LoginParam
	if err := c.Bind(&prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	errors := ValidationResult{}
	if err := c.Validate(prm); err != nil {
		errors = ReportValidation(err)
	}

	user, ok, err := h.db.VerifyUser(prm.Login, prm.Password)
	if err != nil {
		return err
	}

	// success login
	if ok {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   SessionMaxAge,
			HttpOnly: true,
		}
		sess.Values[SessionKeyUserId] = user.ID
		sess.Save(c.Request(), c.Response())
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.Render(http.StatusOK, TmplLogin, LoginView{
		Errors: errors,
	})
}

// GetLogout ログアウト処理
func (h Handler) GetLogout(c echo.Context) error {
	return c.Render(http.StatusOK, TmplLogout, nil)
}
