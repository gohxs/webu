package webu

import (
	"net/http"
	"strings"
)

//Params fetches extra path from handler and return as an array
func Params(r *http.Request) []string {
	server := r.Context().Value(http.ServerContextKey).(*http.Server)
	mux := server.Handler.(*http.ServeMux)
	_, pth := mux.Handler(r)
	extra := strings.TrimPrefix(r.URL.Path, pth)
	if extra == "" {
		return []string{}
	}
	parts := strings.Split(extra, "/")

	return parts
}
