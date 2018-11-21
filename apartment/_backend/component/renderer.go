package component

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo"
)

type rendererOption func(tmpl *templateRenderer) error

func RendererDevMode(dev bool) rendererOption {
	return func(tmpl *templateRenderer) error {
		tmpl.dev = dev
		return nil
	}
}

func NewTemplateRenderer(opts ...rendererOption) (*templateRenderer, error) {
	tmpls, err := loadTemplate()
	if err != nil {
		return nil, err
	}

	renderer := &templateRenderer{
		templates: tmpls,
	}

	for _, opt := range opts {
		err := opt(renderer)
		if err != nil {
			return nil, err
		}
	}

	return renderer, nil
}

func loadTemplate() (map[string]*template.Template, error) {
	files, err := filepath.Glob("./template/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to find template: %v", err)
	}

	tmpls := make(map[string]*template.Template)
	for _, file := range files {
		_, name := filepath.Split(file)
		tmplFiles := []string{
			"./template/shared/base.tmpl",
			file,
		}
		tmpls[name] = template.Must(template.ParseFiles(tmplFiles...))
	}
	return tmpls, nil
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type templateRenderer struct {
	dev       bool
	templates map[string]*template.Template
}

// Render renders a template document
func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// DEV mode
	if t.dev {
		tmpls, err := loadTemplate()
		if err != nil {
			return err
		}
		t.templates = tmpls
	}

	template, ok := t.templates[name]
	if !ok {
		return errors.New("failed to get template")
	}
	// Add global methods if data is a map
	// if viewContext, isMap := data.(map[string]interface{}); isMap {
	//     viewContext["reverse"] = c.Echo().Reverse
	// }

	return template.ExecuteTemplate(w, "base", data)
}
