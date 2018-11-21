package http

import (
	"github.com/Mushus/app/stdapp/server/middleware"
)

// HandlerFuncWithSession 定義されたセッションを持ったhttpのエンドポイントです
type HandlerFuncWithSession func(middleware.SessCtx) error

// HandlerWithSession echo の標準handlerを独自のセッションを持ったhandlerに変更します
func HandlerWithSession(handler HandlerFuncWithSession) echo.HandlerFunc {
	return func(c echo.Context) error {
		sc, ok := c.(*middleware.SessCtx)
		if !ok {
			panic("invalid HandlerWithSession useage")
		}
		return handler(sc)
	}
}
