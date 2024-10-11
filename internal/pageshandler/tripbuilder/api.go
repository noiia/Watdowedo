package tripbuilder

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetFormData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok ok")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	resp, err := http.PostForm("http://www.watdowedo.local/tripbuilder/form", url.Values{"City": {"Values"}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(resp)
	return
}
