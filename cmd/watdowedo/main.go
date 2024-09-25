package main

import (
	"flag"
	"net/http"
	"path/filepath"

	"watdowedo/internal/common/logger"
	"watdowedo/internal/pageshandler/home"
	"watdowedo/internal/pageshandler/tripbuilder"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output in terminal")

	flag.Parse()

	logger, err := logger.Logger(*verbose)
	if err != nil {
		panic(err)
	}

	assetsPath := filepath.Join("web", "static")
	fs := http.FileServer(http.Dir(assetsPath))

	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("GET /home", home.HomeHandler)

	router.HandleFunc("GET /tripbuilder", tripbuilder.TripBuilderHandler)
	router.HandleFunc("POST /tripbuilder/form", tripbuilder.GetFormData)

	logger.Info("starting server at http://localhost" + server.Addr + "/home")

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server internal error : " + err.Error())
	}

}
