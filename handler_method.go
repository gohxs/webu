package webu

import "net/http"

type MethodHandler map[string]http.HandlerFunc

func (m MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m[r.Method]
	if ok {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed\r\n"))

}
