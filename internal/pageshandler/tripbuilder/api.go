package tripbuilder

import (
	"fmt"
	"io"
	"net/http"
)

func GetFormContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	fmt.Println("Données reçues :", string(body))

	fmt.Fprintf(w, "Données reçues avec succès")
}
