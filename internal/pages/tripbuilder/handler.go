package tripbuilder

import (
	"net/http"
	"watdowedo/internal/common/render"
)

func TripBuilderHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "tripbuilder")
}
