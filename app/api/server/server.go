package server

import (
	"fmt"

	"github.com/Mushus/app/api/server/http/endpoint"
	"github.com/Mushus/app/api/server/http/server"

	"github.com/Mushus/app/api/config"
	"github.com/labstack/echo"
)

// Server はサーバーです
type Server struct {
	router   *echo.Echo
	port     int
	entities config.Model
	cfg      config.Config
}

// NewServer はサーバーを生成します
func NewServer(cfg config.Config) Server {
	e := echo.New()
	// NOTE: echoはデフォルトでバナー等が表示されるが不要なので消す
	e.HideBanner = true
	e.HidePort = true

	return Server{
		router: e,
		port:   cfg.App.Port,
		cfg:    cfg,
	}
}

// Start はサーバーを開始してhttpリクエストを受け付ける状態にします
func (s Server) Start() error {
	s.initRouter()
	addr := s.addr()
	if err := s.router.Start(addr); err != nil {
		s.router.Logger.Fatal(err)
		return err
	}
	return nil
}

// アドレスを取得します
func (s Server) addr() string {
	return fmt.Sprintf(":%d", s.port)
}

// ルーターを初期化します
func (s Server) initRouter() {
	router := s.router
	// エンドポイントの設定
	for _, model := range s.models {
		ep := endpoint.NewHandler(model.Key, model)
		ep.AttachRoute(router)
	}
	sh := server.NewHandler()
	sh.AttachRoute(router)
}
