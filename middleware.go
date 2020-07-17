package middleware

import (
	"net/http"
)

func Nest(handlers []http.HandlerFunc) http.HandlerFunc {
	//var result http.HandlerFunc = nil
	//for i := len(handlers) - 1; i > 0; i-- {
	//	if i < len(handlers) - 1 {
	//		result = handlers[i]
	//	}
	//}
	return func(writer http.ResponseWriter, request *http.Request) {
		for _, handler := range handlers {
			handler.ServeHTTP(writer, request)
		}
	}
}

