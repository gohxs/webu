package webu

import (
	"context"
	"net/http"
	"strings"
)

type paramKeyType struct{}

var (
	paramKey = paramKeyType{}
)

// ParamHandler handles and parameterize the extra, and place it in context
func ParamHandler(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		server := r.Context().Value(http.ServerContextKey).(*http.Server)
		mux := server.Handler.(*http.ServeMux)
		_, pth := mux.Handler(r)
		extra := strings.TrimPrefix(r.URL.String(), pth)
		parts := strings.Split(extra, "/")

		nr := r.WithContext(context.WithValue(r.Context(), paramKey, parts))

		next(w, nr)
	}
}

// Param from request context
func Param(r *http.Request) []string {
	if params, ok := r.Context().Value(paramKey).([]string); ok {
		return params
	}
	return nil

}
