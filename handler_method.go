package webu

import "net/http"

//Method router
type Method map[string]http.HandlerFunc

func (m Method) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m[r.Method]
	if ok {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed\r\n"))
}

// MethodHandler returns a func
func MethodHandler(method string, handlerFunc http.HandlerFunc) http.HandlerFunc {

	return Method{method: handlerFunc}.ServeHTTP
}
