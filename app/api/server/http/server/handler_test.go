package server_test

import (
	"testing"

	"github.com/Mushus/app/api/server/http/server"
	"github.com/labstack/echo"
)

func TestDuckType(t *testing.T) {
	h := server.NewHandler()
	var _ echo.HandlerFunc = h.HealthCheck
}
