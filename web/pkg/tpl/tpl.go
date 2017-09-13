package tpl

import (
	"errors"
	"html/template"

	"github.com/gocraft/web"
)

// Data struct
type Data struct {
	TemplateFile string
	Data         interface{}
}

// Render func
func (s Data) Render(w web.ResponseWriter, r *web.Request) error {
	if s.TemplateFile == "" {
		tplErr := errors.New("tpl.Render requires non-empty TemplateFile")
		return tplErr
	}

	t := template.Must(template.ParseFiles("templates/"+s.TemplateFile, "templates/layout.html"))
	err := t.ExecuteTemplate(w, "base", s)

	return err
}
