package app

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/middleware"
)

// Manager はアプリケーションの管理を行います
type Manager struct {
	apps   map[string]Application
	client *http.Client
}

// Application はアプリケーションを表します
type Application struct {
	URL url.URL
}

// NewManager アプリケーションマネージャーを作成する
func NewManager() *Manager {
	return &Manager{
		apps:   map[string]Application{},
		client: &http.Client{},
	}
}

// Register はアプリケーションを登録します
func (m Manager) Register(key string, app Application) {
	m.apps[key] = app
}

func (m Manager) getApp(key string) (Application, bool) {
	app, ok := m.apps[key]
	return app, ok
}

// Serve proxyします
func (m Manager) Serve(key string, ctx context.Context, path string, req *http.Request) error {
	app, ok := m.getApp(key)
	if !ok {
		return errors.New("appication not found")
	}
	middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: &app.URL,
			},
		}),
	})
	// appReq := http.NewRequest(req.Method, , req.Body)
	// m.client.Do()
	return nil
}
