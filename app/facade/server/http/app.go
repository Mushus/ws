package http

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mushus/app/facade/app"
	"github.com/labstack/echo"
)

// NewAppHandler は AppHandler を作成します
func NewAppHandler(am *app.Manager, key string) AppHandler {
	return AppHandler{
		appManager: am,
		key:        key,
	}
}

// AppHandler はアプリケーションとゲートウェイの接続するエンドポイントです
type AppHandler struct {
	appManager *app.Manager
	key        string
}

// PathPrefix はpathのprefixを取得します
func (a AppHandler) PathPrefix() string {
	return fmt.Sprintf("/%s", a.key)
}

// Handle はhttpリクエストを受け付けます
func (a AppHandler) Handle(c echo.Context) error {
	req := c.Request()
	ctx := context.Background()
	path := strings.Trim(req.RequestURI, a.PathPrefix())
	return a.appManager.Serve(a.key, ctx, path, req)
}
