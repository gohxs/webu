package webu

import (
	"net/http"
)

//Methods router
//  mux.Handle("/api/v1/user", Methods{"GET": handleGet, "POST", handlePost})
type Methods map[string]http.HandlerFunc

//ServeHTTP default handler implementation
func (m Methods) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m[r.Method]
	if ok {
		handler(w, r)
		return
	}
	WriteStatus(w, http.StatusMethodNotAllowed)
}

// Get handles Get Method
func (m *Methods) Get(handlerFunc http.HandlerFunc) {
	(*m)["GET"] = handlerFunc
}

// Post handles Get Method
func (m *Methods) Post(handlerFunc http.HandlerFunc) {
	(*m)["POST"] = handlerFunc
}

//MethodHandler returns a func
//Single method handler, cannot be chained since it returns not allowed if other methods are called
// mux.HandleFunc("/api/v1/user",webu.MethodHandler("GET",handleGet))
func MethodHandler(method string, handlerFunc http.HandlerFunc) http.HandlerFunc {

	return Methods{method: handlerFunc}.ServeHTTP
}
