package webu

import (
	"net/http"
)

// MapHandler serves a []byte request based on Request
func MapHandler(m map[string][]byte, catch interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.String()[1:] // skip '/'
		d, ok := m[urlPath]
		if !ok {
			switch t := catch.(type) {
			case http.HandlerFunc:
				t(w, r) // catchHandler
			case string:
				w.Write(m[t])
			}
		}
		w.Write(d)
	})
}
