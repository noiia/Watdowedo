package tripbuilder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"watdowedo/internal/common/logger"
)

type Data_Structure struct {
	Destination  string `json:"destination"`
	WalkingLevel string `json:"walking-level"`
	Validity     string `json:"validity"`
	Drive        string `json:"drive"`
}

func GetFormData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "./tripbuilder")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	decoder := json.NewDecoder(r.Body)

	var unmarshaledValues Data_Structure

	if err := decoder.Decode(&unmarshaledValues); err != nil {
		logger.GlobalLogger.Error("decoding json error from http://watdowedo : " + r.URL.Path + " : " + err.Error())
	}

	fmt.Println(unmarshaledValues)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w)

	return
}
