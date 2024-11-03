package trip

import (
	"errors"
	"fmt"
	"net/http"
	"watdowedo/internal/common/logger"
)

func getID(r *http.Request) (id string, err error) {
	for i := range len(r.URL.Path) {
		index := len(r.URL.Path) - (i + 1)
		if string(r.URL.Path[index]) != "" {
			if string(r.URL.Path[index]) == "/" {
				return id, nil
			} else {
				id = string(r.URL.Path[index])
			}
		} else {
			return id, errors.New("no url path found")
		}
	}
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var id string
	var err error

	if id, err = getID(r); err != nil {
		logger.GlobalLogger.Error(r.Method + " - " + err.Error() + " : at " + r.URL.Path)
	}

	fmt.Println("id :" + id)
}
