package gomuxrender

import (
	"html/template"
	"net/http"
)

type MuxRenderer struct {
	templates *template.Template
}

// NewMuxRenderer creates a MuxRenderer witch contains parsed templates found in the
// directory that matches the pattern (1º param).
//
// "Under the hood" it calls template.Must(template.ParseGlob(pattern)) and returns
// the templates to the MuxRenderer.
func NewMuxRenderer(pattern string) *MuxRenderer {
	return &MuxRenderer{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

// Render recieves the http.ResponseWriter (1º Param) and executes the template
// whose name matches with name (2º param) towards it, passing down data contained
// in data (3º param).
func (mr *MuxRenderer) Render(w http.ResponseWriter, name string, data interface{}) {
	mr.templates.ExecuteTemplate(w, name, data)
}
