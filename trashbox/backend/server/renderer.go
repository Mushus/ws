package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type (
	// TemplateMap 複数のテンプレートをまとめたもの
	TemplateMap map[string]*template.Template
	// Renderer テンプレートエンジン
	Renderer struct {
		templates TemplateMap
	}
)

// ExecuteTemplate テンプレートを実行する
func (t TemplateMap) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	if tmpl, ok := t[name]; ok {
		return tmpl.ExecuteTemplate(w, "layout", data)
	}
	return xerrors.Errorf("template %v is not found", name)
}

// NewRenderer 新しいテンプレートを作成する
func NewRenderer() *Renderer {
	return &Renderer{
		templates: templates,
	}
}

func composeTemplate(content *template.Template) *template.Template {
	tmpl := template.Must(layoutTmpl.Clone())
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

const (
	// TmplLogin ログイン
	TmplLogin string = "login"
	// TmplLogout ログアウト
	TmplLogout string = "logout"
)

var templates = TemplateMap{
	TmplLogin:  composeTemplate(loginTmpl),
	TmplLogout: composeTemplate(logoutTmpl),
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

var loginTmpl = template.Must(template.New("").Parse(`
{{define "title"}}Login{{end}}
{{define "content"}}
{{range .Errors}}
<div class="errors">
{{range .}}{{.}}{{end}}</div>
{{end}}
<form method="POST" action="login">
<input type="text" name="login" placeholder="login_name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
{{end}}
`))

var logoutTmpl = template.Must(template.New("").Parse(`
{{define "title"}}Logout{{end}}
{{define "content"}}
<p>ログアウトしました</p>
{{end}}
`))
