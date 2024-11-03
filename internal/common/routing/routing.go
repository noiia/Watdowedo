package routing

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"

	"watdowedo/internal/common/logger"
	"watdowedo/internal/common/render"
	"watdowedo/internal/pageshandler/trip"
	"watdowedo/internal/pageshandler/tripbuilder"
)

type routeType struct {
	method   string
	path     *regexp.Regexp
	filename string
	handler  http.HandlerFunc
}

func newRoute(method string, pattern string, filename string, handler http.HandlerFunc) routeType {
	return routeType{method, regexp.MustCompile(pattern), filename, handler}
}

func commonHandler(rt routeType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.RenderTemplates(w, rt.filename)
	}
}

var routes = []routeType{
	newRoute("GET", "/", "home", commonHandler(routeType{filename: "home"})),

	newRoute("GET", "/tripbuilder", "tripbuilder", commonHandler(routeType{filename: "tripbuilder"})),
	newRoute("POST", "/tripbuilder/form", "", tripbuilder.GetFormData),

	newRoute("GET", "/login", "login", commonHandler(routeType{filename: "login"})),
	newRoute("POST", "/login/form", "", tripbuilder.GetFormData),

	newRoute("GET", "/forgottenpw", "forgottenpw", commonHandler(routeType{filename: "forgottenpw"})),
	newRoute("POST", "/forgottenpw/form", "", tripbuilder.GetFormData),

	newRoute("GET", "/trip/", "trip", trip.Handler),
}

func Routing() http.Server {
	assetsPath := filepath.Join("web", "static")
	fs := http.FileServer(http.Dir(assetsPath))

	router := http.NewServeMux()

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	for _, route := range routes {
		func(rt routeType) {
			router.HandleFunc(rt.path.String(), func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(r.URL.Path, rt.path.String())
				matches := route.path.FindStringSubmatch(r.URL.Path)
				if len(matches) == 0 {
					http.Error(w, "bad path", http.StatusNotFound)
					logger.GlobalLogger.Error(r.Method + strconv.Itoa(http.StatusNotFound) + " : bad path " + r.URL.Path)
					return
				}

				if r.Method != rt.method {
					http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
					logger.GlobalLogger.Error(r.Method + strconv.Itoa(http.StatusMethodNotAllowed) + " : Method not allowed at http://watdowedo" + r.URL.Path)
					return
				}

				logger.GlobalLogger.Info(r.Method + " : http://watdowedo" + r.URL.Path)

				rt.handler(w, r)
			})
		}(route)
	}

	return http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
