package server

import (
	"net/http"
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
func (h Handler) GetLogin(c Context) error {
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
func (h Handler) PostLogin(c Context) error {
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
		sess, _ := getSession(c)
		sess.Values[SessionKeyUserID] = user.ID
		if err := saveSession(c, sess); err != nil {
			return err
		}
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.Render(http.StatusOK, TmplLogin, LoginView{
		Errors: errors,
	})
}

// GetLogout is a handler to logout users
func (h Handler) GetLogout(c Context) error {
	// TODO: logout process
	return c.Render(http.StatusOK, TmplLogout, nil)
}

// GetIndex is a handler show index of webpage
func (h Handler) GetIndex(c Context) error {
	return c.String(http.StatusOK, "it's works!")
}

// GetDoc is a handler of get document
func (h Handler) GetDoc(c Context) error {
	name := c.Param("name")

	doc, err := h.docs.Get(name)
	if err == DocumentNotFound {
		if !c.IsLoggedIn {
			return c.String(http.StatusNotFound, "document not found")
		}
		return c.Render(http.StatusOK, TmplEdit, nil)
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, doc.Body)
}
