package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type Index struct {
}

func (h Index) RedirectToIndex(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/admin/")
}

func (h Index) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "app.tmpl", map[string]interface{}{})
}
