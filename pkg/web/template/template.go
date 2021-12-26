package template

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateUI struct {
	Templates *template.Template
}

func (t *TemplateUI) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
