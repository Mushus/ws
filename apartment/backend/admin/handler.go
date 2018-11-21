package admin

import (
	"bytes"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/Mushus/apartment/backend/gql"
	"github.com/Mushus/apartment/backend/service"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func NewSessionController(db *sqlx.DB, auth service.AuthService) *sessionController {
	return &sessionController{
		DB:   db,
		Auth: auth,
	}
}

// SessionController
type sessionController struct {
	DB   *sqlx.DB
	Auth service.AuthService
}

func (s sessionController) Login(c echo.Context) error {
	prm := &LoginParam{}
	if err := c.Bind(prm); err != nil {
		return c.JSON(http.StatusBadRequest, BadRequest)
	}
	sql, args, err := sq.Select("*").From("users").Where("login = ?", prm.Login).ToSql()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, InternalServerError)
	}

	admin := new(Admin)
	if err := s.DB.Get(admin, sql, args...); err != nil {
		return c.JSON(http.StatusInternalServerError, InternalServerError)
	}

	if !admin.ComparePassword(prm.Password) {
		c.JSON(http.StatusUnauthorized, Unauthorized)
	}

}

func (s sessionController) Logout(c echo.Context) error {
	return nil
}

func NewGrahpQLController(db *sqlx.DB) *grahpQLController {
	schema := gql.NewSchema(db)
	return &grahpQLController{
		DB:     db,
		Schema: schema,
	}
}

type grahpQLController struct {
	DB     *sqlx.DB
	Schema graphql.Schema
}

func (g grahpQLController) GraphQL(c echo.Context) error {
	bufBody := &bytes.Buffer{}
	bufBody.ReadFrom(c.Request().Body)
	params := graphql.Params{
		Schema:        g.Schema,
		RequestString: bufBody.String(),
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		return fmt.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return c.JSON(http.StatusOK, r)
}
