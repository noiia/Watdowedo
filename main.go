package main

import (
	"fmt"
	"net/http"

	"watdowedo/internal/pages/home"
	"watdowedo/internal/pages/tripbuilder"
)

const SERVER_PORT = ":8080"

func main() {
	http.HandleFunc("/home", home.HomeHandler)
	http.HandleFunc("/trip-builder", tripbuilder.TripBuilderHandler)

	fmt.Println("(http://localhost" + SERVER_PORT + ")")

	http.ListenAndServe(SERVER_PORT, nil)
}
