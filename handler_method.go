package webu

import "net/http"

// Handle multiple methods
type MethodHandler struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Put    http.HandlerFunc
	Delete http.HandlerFunc
}

func (m *MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	helperMap := map[string]http.HandlerFunc{
		"GET": m.Get, "POST": m.Post,
		"PUT": m.Put, "DELETE": m.Delete,
	}

	handler, ok := helperMap[r.Method]
	if ok {
		handler(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed\r\n"))
}
