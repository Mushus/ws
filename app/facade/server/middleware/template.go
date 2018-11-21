package middleware

import (
	"html/template"
	"io"
	"path"
	"path/filepath"

	"github.com/labstack/echo"
)

const (
	// TODO: 設定ファイル化
	ViewDir = "./view"
)

func NewRenderer() (TemplateRenderer, error) {
	fp := filepath.Join(ViewDir, "template", "*.tmpl")
	// HACK: Mustやめる
	return TemplateRenderer{
		shared: template.Must(template.ParseGlob(fp)),
	}, nil
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	shared    *template.Template
	templates []template.Template
}

// Render renders a template document
func (t TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// 追加する値等
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	filename := name + ".tmpl"
	fp := filepath.Join(ViewDir, "pages", path.Clean("/"+filename)) // NOTE: セキュリティ対策
	shared, err := t.shared.Clone()
	if err != nil {
		return err
	}
	shared, err = shared.ParseFiles(fp)
	if err != nil {
		return err
	}
	return shared.ExecuteTemplate(w, "template.tmpl", data)
}
