package main

import (
	"flag"
	"fmt"
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

	assetsPath := filepath.Join("web", "templates")
	fs := http.FileServer(http.Dir(assetsPath))
	http.Handle("/"+assetsPath+"/", http.StripPrefix("/"+assetsPath, fs))

	http.HandleFunc("/home", home.HomeHandler)
	http.HandleFunc("/trip-builder", tripbuilder.TripBuilderHandler)

	fmt.Println("(http://localhost" + SERVER_PORT + "/home)")
	logger.Info("starting server at http://localhost" + SERVER_PORT + "/home")

	http.ListenAndServe(SERVER_PORT, nil)
}
