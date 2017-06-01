package webu

import "net/http"

type Routes []Entry

type Entry struct {
	Path     string
	Methods  []string
	Chain    *Chain
	Meta     map[string]interface{} // Extra info for certain things?
	Handler  http.HandlerFunc
	Children Routes
}
