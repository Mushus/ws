package endpoint_test

import (
	"testing"

	"github.com/Mushus/app/api/config"
	http "github.com/Mushus/app/api/server/http/endpoint"
	"github.com/labstack/echo"
)

func TestDuckType(t *testing.T) {
	model := config.Model{
		Key: "ok",
	}

	e := http.NewHandler(model)
	var _ echo.HandlerFunc = e.Create
	var _ echo.HandlerFunc = e.Update
	var _ echo.HandlerFunc = e.Read
	var _ echo.HandlerFunc = e.Delete
}
