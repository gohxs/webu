package webu

import (
	"net/http"

	"github.com/gohxs/webu/chain"
)

type Routes []Entry

type Entry struct {
	Path     string
	Methods  []string
	Chain    *chain.Chain
	Meta     map[string]interface{} // Extra info for certain things?
	Handler  http.HandlerFunc
	Children Routes
}
