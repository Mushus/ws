package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type Frontend struct {
}

func (h Frontend) Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
