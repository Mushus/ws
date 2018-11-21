package http

import (
	"html/template"

	"github.com/labstack/echo"
)

type Dashboard struct {
}

func NewDashboard() Dashboard {
	return Dashboard{}
}

func (d Dashboard) Handle(c echo.Context) error {
	return c.Render(200, "dashboard", View{
		Head: ViewHead{
			Title: "ダッシュボード",
			Scripts: []ViewScript{
				{Src: "/static/js/main.js"},
				{Src: "/static/js/component.js"},
			},
			WidgetScripts: []ViewWidgetScripts{
				{Src: "/static/js/component/widget1.js", TagName: "simple-dashboard"},
			},
		},
		Body: ViewBody{
			Nav: []ViewNav{
				{
					Title: "コンテンツの管理",
					Submenu: []ViewNavLink{
						{Link: "/text", Text: "hoge"},
					},
				},
			},
			Main: ViewDashboard{
				Dashboard: []ViewDWidget{
					{
						Title: "widget1",
						Tag:   template.HTML("<simple-dashboard></simple-dashboard>"),
					},
				},
			},
		},
	})
}

type ViewDashboard struct {
	Dashboard []ViewDWidget
}

type ViewDWidget struct {
	Title string
	Tag   template.HTML
}
