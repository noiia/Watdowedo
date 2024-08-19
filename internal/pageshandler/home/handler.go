package home

import (
	"net/http"
	"watdowedo/internal/common/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home")
}
