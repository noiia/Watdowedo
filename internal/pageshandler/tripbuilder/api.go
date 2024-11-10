package tripbuilder

import (
	"encoding/json"
	"net/http"
	"strings"
	"watdowedo/internal/common/logger"
)

type Trip_builder_form struct {
	Destination  string `json:"destination"`
	WalkingLevel string `json:"walking-level"`
	Validity     string `json:"validity"`
	Drive        string `json:"drive"`
}

func (w Trip_builder_form) string() string {
	return strings.Join([]string{w.Destination, w.WalkingLevel, w.Validity, w.Drive}, " ")
}

func (w Trip_builder_form) complete() bool {
	fieldList := []string{w.Destination, w.WalkingLevel, w.Validity, w.Drive}
	for _, value := range fieldList {
		if value == "" {
			return false
		}
	}

	return true
}

func GetFormData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "./tripbuilder")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	decoder := json.NewDecoder(r.Body)

	var unmarshaledValues Trip_builder_form

	if err := decoder.Decode(&unmarshaledValues); err != nil {
		logger.GlobalLogger.Error("decoding json error from http://watdowedo : " + r.URL.Path + " : " + err.Error())
	}

	logger.GlobalLogger.Info(unmarshaledValues.string())

	if unmarshaledValues.complete() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w)
	}

	return
}
