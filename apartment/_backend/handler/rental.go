package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type Rental struct {
}

func (h Rental) Print(c echo.Context) error {
	return c.Render(http.StatusOK, "print.tmpl", map[string]interface{}{})
}
