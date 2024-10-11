package routing

import (
	"net/http"
	"path/filepath"
	"strconv"

	"watdowedo/internal/common/logger"
	"watdowedo/internal/pageshandler/home"
	"watdowedo/internal/pageshandler/tripbuilder"
)

var routes = []routeType{
	newRoute("GET", "/", home.HomeHandler),

	newRoute("GET", "/tripbuilder", tripbuilder.Handler),
	newRoute("POST", "/tripbuilder/form", tripbuilder.GetFormData),

	newRoute("GET", "/login", tripbuilder.Handler),
	newRoute("POST", "/login/form", tripbuilder.GetFormData),

	newRoute("GET", "/forgottenpw", tripbuilder.Handler),
	newRoute("POST", "/forgottenpw/form", tripbuilder.GetFormData),
}

func newRoute(method, pattern string, handler http.HandlerFunc) routeType {
	return routeType{method, pattern, handler}
}

type routeType struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func Routing() http.Server {
	assetsPath := filepath.Join("web", "static")
	fs := http.FileServer(http.Dir(assetsPath))

	router := http.NewServeMux()

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	for _, route := range routes {
		func(rt routeType) {
			router.HandleFunc(rt.path, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != rt.path {
					http.Error(w, "bad path", http.StatusMethodNotAllowed)
					logger.GlobalLogger.Error(strconv.Itoa(http.StatusMethodNotAllowed) + " : bad path " + r.URL.Path)
					return
				}

				if r.Method != rt.method {
					http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
					logger.GlobalLogger.Error(strconv.Itoa(http.StatusMethodNotAllowed) + " : Method not allowed at https://watdowedo" + r.URL.Path)
					return
				}

				rt.handler(w, r)
			})
		}(route)
	}

	return http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
