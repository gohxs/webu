package webu

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type SHttpFunc func(...string) (interface{}, error)

// Transform
func SpecialHandler(prefix string, sfunc SHttpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received:", r.URL.Path, " Chek:", prefix)
		spath := strings.TrimPrefix(r.URL.Path, prefix)
		log.Println("Result:", spath)

		params := strings.Split(spath, "/")

		obj, err := sfunc(params...)
		if err != nil {
			log.Println("Special func error")
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}
