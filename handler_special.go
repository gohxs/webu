// Experiment not stable
package webu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type SpecialBase struct {
	W http.ResponseWriter
	R *http.Request
}

type SpecialInterface interface{}
type SpecialCreator func(*SpecialBase) SpecialInterface

type SHttpFunc func(...string) (interface{}, error)

// We will build a new interface and execute with Exec

func SpecialHandler(prefix string, create SpecialCreator, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj := create(&SpecialBase{w, r})
		method := reflect.ValueOf(obj).MethodByName(name).Interface().(func(...string) (interface{}, error))

		handler := SpecialHandlerFunc(prefix, method)

		handler(w, r)
	})
}

// SpecialHandler Transform default HandlerFunc to our SHttpFunc
func SpecialHandlerFunc(prefix string, sfunc SHttpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		spath := strings.TrimPrefix(r.URL.Path, prefix)
		spath = strings.Trim(spath, "/")

		params := strings.Split(spath, "/")

		obj, err := sfunc(params...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintln("Error:", err)))
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}
