package render

import (
	"html/template"
	"net/http"
	"path/filepath"
	basetemplate "watdowedo/web/templates"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	template, err := template.ParseFiles(filepath.Join("web", "templates", tmpl+".page.tmpl"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	base := basetemplate.LoadBase()
	if err := template.Execute(w, base); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
