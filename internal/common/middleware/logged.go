package middleware

import "net/http"

func IsLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		if user == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
