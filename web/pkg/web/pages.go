package web

import (
	"github.com/wizelineacademy/GoWorkshop/web/pkg/tpl"

	"github.com/gocraft/web"
)

func (c *Context) home(w web.ResponseWriter, r *web.Request) {
	d := tpl.Data{
		TemplateFile: "pages/home.html",
	}

	d.Render(w, r)
}
