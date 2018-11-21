package server

import "github.com/labstack/echo"

// Handler ヘルスチェックを行うエンドポイントです
type Handler struct {
}

// NewHandler はインスタンスを生成します。
func NewHandler() Handler {
	return Handler{}
}

// AttachRoute ルーターを設定します
func (h Handler) AttachRoute(router *echo.Echo) {
	g := router.Group("/.server")
	g.GET("/health_check", h.HealthCheck)
}

// HealthCheck はヘルスチェックを実行します
func (h Handler) HealthCheck(c echo.Context) error {
	resp := Response{
		Status: StatusOK,
	}
	return c.JSON(200, resp)
}
