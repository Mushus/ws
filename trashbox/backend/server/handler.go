package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Handler handler
type Handler struct {
	db   *DB
	docs *DocRepo
}

// NewHandler ハンドラを生成する
func NewHandler(db *DB, docs *DocRepo) Handler {
	return Handler{
		db:   db,
		docs: docs,
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
		sess.Values[SessionKeyUserID] = user.ID
		sess.Save(c.Request(), c.Response())
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.Render(http.StatusOK, TmplLogin, LoginView{
		Errors: errors,
	})
}

// GetLogout is a handler to logout users
func (h Handler) GetLogout(c echo.Context) error {
	// TODO: logout process
	return c.Render(http.StatusOK, TmplLogout, nil)
}

// GetIndex is a handler show index of webpage
func (h Handler) GetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "it's works!")
}

// GetDoc is a handler of get document
func (h Handler) GetDoc(c echo.Context) error {
	name := c.Param("name")

	doc, err := h.docs.Get(name)
	if err == DocumentNotFound {
		return c.String(http.StatusNotFound, "404 document not found")
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, doc.Body)
}
