package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

// TemplateMap 複数のテンプレートをまとめたもの
type TemplateMap map[string]*template.Template

// ExecuteTemplate テンプレートを実行する
func (t TemplateMap) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	if tmpl, ok := t[name]; ok {
		return tmpl.ExecuteTemplate(w, "layout", data)
	}
	return xerrors.Errorf("template %v is not found", name)
}

// Renderer テンプレートエンジン
type Renderer struct {
	templates TemplateMap
}

// NewRenderer 新しいテンプレートを作成する
func NewRenderer() *Renderer {
	templates := TemplateMap{
		"login": composeTemplate(layoutTmpl, loginTmpl),
	}
	return &Renderer{
		templates: templates,
	}
}

func composeTemplate(layout *template.Template, content *template.Template) *template.Template {
	tmpl := template.Must(layout.Clone())
	titleBody := content.Lookup("title")
	tmpl = template.Must(tmpl.AddParseTree("title", titleBody.Tree))
	contentBody := content.Lookup("content")
	tmpl = template.Must(tmpl.AddParseTree("content", contentBody.Tree))
	return tmpl
}

// Render レンダリングを行います
func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var layoutTmpl = template.Must(template.New("layout").Parse(`<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>{{template "title" .}}</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
<body>
{{template "content" .}}
</body>
</html>
`))

// LoginView ログインのビューモデル
type LoginView struct {
	Errors ValidationResult
}

var loginTmpl = template.Must(template.New("login").Parse(`
{{define "title"}}Login{{end}}
{{define "content"}}
{{range .Errors}}
<div class="errors">
{{range .}}{{.}}{{end}}</div>
{{end}}
<form method="POST" action="login">
<input type="text" name="user" placeholder="user name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
{{end}}
`))
