package main

import (
	"fmt"
	"net/http"

	"watdowedo/internal/pages/home"
)

const SERVER_PORT = ":8080"

func main() {
	http.HandleFunc("/", home.Home)
	fmt.Println("(http://localhost" + SERVER_PORT + ")")
	http.ListenAndServe(SERVER_PORT, nil)
}
