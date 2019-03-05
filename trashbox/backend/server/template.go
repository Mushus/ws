package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template テンプレートエンジン
type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: loginTemplate,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var loginTemplate = template.Must(template.New("login").Parse(`<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
<body>
<form method="POST" action="login">
<input type="text" name="user" placeholder="user name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
</body>
</html>
`))
