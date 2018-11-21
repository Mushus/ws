package server

import (
	"github.com/Mushus/app/facade/server/middleware"
	"fmt"
	"net/url"

	"github.com/Mushus/app/facade/app"
	"github.com/Mushus/app/facade/config"
	"github.com/Mushus/app/facade/server/http"
	"github.com/labstack/echo"
)

const (
	staticPath = "./static"
)

// Server はサーバーです
type Server struct {
	router *echo.Echo
	port   int
	cfg    config.Config
}

// NewServer はサーバーを生成します
func NewServer(cfg config.Config) (*Server, error) {
	renderer , err := middleware.NewRenderer()
	if err != nil {
		return nil, err
	}
	e := echo.New()
	// NOTE: echoはデフォルトでバナー等が表示されるが不要なので消す
	e.HideBanner = true
	e.HidePort = true

	// TODO: デバッグモード
	e.Debug = true
	e.Renderer = renderer

	e.Static("/static", staticPath)

	am, err := initAppManager(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init appManager: %v", err)
	}

	initRouter(e, am, cfg.Applications)

	return &Server{
		router: e,
		port:   cfg.Server.Port,
		cfg:    cfg,
	}, nil
}

// Start はサーバーを開始してhttpリクエストを受け付ける状態にします
func (s Server) Start() error {
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

func initAppManager(cfg config.Config) (*app.Manager, error) {
	am := app.NewManager()
	for key, appcfg := range cfg.Applications {
		url, err := url.Parse(appcfg.URL)
		if err != nil {
			return nil, err
		}
		am.Register(key, app.Application{
			URL: *url,
		})
	}
	return am, nil
}

// ルーターを初期化します
func initRouter(e *echo.Echo, am *app.Manager, appcfg config.Applications) {
	// ダッシュボード
	dashboard := http.NewDashboard()
	e.GET("/", dashboard.Handle)

	for k := range appcfg {
		handler := http.NewAppHandler(am, k)
		path := fmt.Sprintf("%s/*", handler.PathPrefix())
		e.Any(path, handler.Handle)
	}
}
