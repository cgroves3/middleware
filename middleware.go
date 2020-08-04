package middleware

import (
	"net/http"
)

func JoinHandlerFuncs(handlers ...http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		for _, handler := range handlers {
			handler.ServeHTTP(writer, request)
		}
	}
}

func JoinHandlers(handlers ...http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			handler.ServeHTTP(w, r)
		}
	})
}

