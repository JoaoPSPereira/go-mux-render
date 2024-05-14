package gomuxrender

import (
	"html/template"
	"net/http"
)

type Templates struct {
	templates *template.Template
}

// NewMuxRenderer creates a new Renderer with the parsed templates found in the
// directory that matches the pattern.
//
// "Under the hood" it calls template.Must(template.ParseGlob(pattern)) and returns
// the templates to the Renderer.
func NewMuxRenderer(pattern string) *Renderer {
	return &Renderer{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

// Render recieves the http.ResponseWriter (1º Param) and executes the template
// whose name matches with name (2º param) towards it, passing down data contained
// in data (3º param).
func (r *Renderer) Render(w http.ResponseWriter, name string, data interface{}) {
	r.templates.ExecuteTemplate(w, name, data)
}
