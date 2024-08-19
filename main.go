package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"watdowedo/internal/pages/home"
	"watdowedo/internal/pages/tripbuilder"
)

const SERVER_PORT = ":8080"

func main() {
	http.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, filepath.Join("src", "templates", "output.css"))
	})

	http.HandleFunc("/home", home.HomeHandler)
	http.HandleFunc("/trip-builder", tripbuilder.TripBuilderHandler)

	fmt.Println("(http://localhost" + SERVER_PORT + ")")

	http.ListenAndServe(SERVER_PORT, nil)
}
