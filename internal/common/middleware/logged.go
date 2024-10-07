package middleware

import "net/http"

func IsLogged(http.HandlerFunc) bool {
	return true
}
