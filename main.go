package gomuxrender

import (
	"html/template"
	"net/http"
)

type Templates struct {
	templates *template.Template
}

var Renderer *Templates = &Templates{
	templates: template.Must(template.ParseGlob("public/*.html")),
}

func Render(w http.ResponseWriter, name string, data interface{}) {
	Renderer.templates.ExecuteTemplate(w, name, data)
}
