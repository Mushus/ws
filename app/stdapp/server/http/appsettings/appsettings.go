package appsettings

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// Resp Handler のレスポンスです
	Resp struct {
		Widgets Widgets            `json:"widgets"`
		Sidenav map[string]Sidenav `json:"sidenav"`
	}

	// Widgets ウィジェットです
	Widgets struct {
		Dashboard Widget `json:"dashboard"`
	}

	// Widget ウィジェットの詳細設定です
	Widget struct {
		TemplateHTML string `json:"template_html,omitempty"`
		ComponentJS  string `json:"component_js,omitempty"`
	}

	// Sidenav はサイドメニューに表示する内容です
	Sidenav struct {
		Label string
		Path  string
	}
)

// Handler は /.server/.app-settings を処理するエントリーポイントです
func Handler(c echo.Context) error {
	// TODO: キャッシュ対応のためにヘッダーを追加する
	return c.JSON(http.StatusOK,
		Resp{
			Widgets: Widgets{
				Dashboard: Widget{},
			},
			Sidenav: map[string]Sidenav{
				"hoge": Sidenav{
					Label: "hoge",
					Path:  "/hoge",
				},
			},
		},
	)
}
