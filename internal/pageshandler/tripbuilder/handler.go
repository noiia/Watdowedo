package tripbuilder

import (
	"net/http"
	"watdowedo/internal/common/render"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "tripbuilder")
}
