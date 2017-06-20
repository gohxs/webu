package webu

import (
	"net/http"
	"strings"
)

// Param fetches extra path from handler and return as an array
func Param(r *http.Request) []string {
	server := r.Context().Value(http.ServerContextKey).(*http.Server)
	mux := server.Handler.(*http.ServeMux)
	_, pth := mux.Handler(r)
	extra := strings.TrimPrefix(r.URL.String(), pth)
	parts := strings.Split(extra, "/")

	return parts
}
