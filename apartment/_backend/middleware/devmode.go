package middleware

import "github.com/labstack/echo"

func DevMode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Response().Header()
		header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
		header.Set("Pragma", "no-cache")
		return next(c)
	}
}
