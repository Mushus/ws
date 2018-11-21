package endpoint

import (
	"fmt"
	"net/http"

	"github.com/Mushus/app/api/config"
	"github.com/labstack/echo"
)

// Handler は http サーバーにアクセスする時に使用されるパスです
type Handler struct {
	key   string
	model config.Model
}

// NewHandler は key という名前のエントリーポイントを作成します
func NewHandler(model config.Model) Handler {
	return Handler{
		key:   model.Key,
		model: model,
	}
}

// Path エントリーポインのパスです
func (e Handler) Path() string {
	return fmt.Sprintf("/%s", e.key)
}

// PathWithID ID付きのエントリーポイントのパスを返します
func (e Handler) PathWithID() string {
	return fmt.Sprintf("/%s/:id", e.key)
}

// AttachRoute エントリーポイントをルーターに設定します
func (e Handler) AttachRoute(router *echo.Echo) {
	path := e.Path()
	pathWithID := e.PathWithID()
	router.POST(path, e.Create)
	router.PUT(pathWithID, e.Update)
	router.GET(pathWithID, e.Read)
	router.DELETE(pathWithID, e.Delete)
}

// Create は http の POST メソッドに相当します
func (e Handler) Create(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("create %s", e.key))
}

// Update は http の PUT メソッドに相当します
func (e Handler) Update(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("update %s", e.key))
}

// Read は http の GET メソッドに相当します
func (e Handler) Read(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("read %s", e.key))
}

// Delete は http の DELETE メソッドに相当します
func (e Handler) Delete(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("delete %s", e.key))
}
