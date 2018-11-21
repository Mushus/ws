package healthcheck

import (
	"net/http"

	"github.com/labstack/echo"
)


type (
	Resp struct {
		Status Status `json:"status"`
		Services map[string]Status `json:"services"`
	}

	Status string
)

const (
	StatusOK = "ok"
	StatusNG = "ng"
)

// HealthCheck ヘルスチェックを行うエンドポイントです
func NewHanlder() echo.HandlerFunc {
	return func Handler(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			Resp {
				Status: StatusOK,
				Services: map[string]string {
					"db": StatusOK,
				}
			}
		)
	}
}
