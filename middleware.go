package middleware

import (
	"net/http"
)

func Join(handlers []http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		for _, handler := range handlers {
			handler.ServeHTTP(writer, request)
		}
	}
}

