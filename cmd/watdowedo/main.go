package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"watdowedo/internal/common/logger"
	"watdowedo/internal/pageshandler/home"
	"watdowedo/internal/pageshandler/tripbuilder"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	verbose := flag.Bool("v", false, "verbose output in terminal")

	flag.Parse()

	logger, err := logger.Logger(*verbose)
	if err != nil {
		panic(err)
	}

	SERVER_PORT := os.Getenv("LISTEN_ADDR")

	assetsPath := filepath.Join("web", "static")
	fs := http.FileServer(http.Dir(assetsPath))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", home.HomeHandler)
	http.HandleFunc("/trip-builder", tripbuilder.TripBuilderHandler)

	logger.Info("starting server at http://localhost" + SERVER_PORT + "/home")

	http.ListenAndServe(SERVER_PORT, nil)
}
