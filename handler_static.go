package webu

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//StaticHandler serves file or execute handler if file not found
// webu.StaticHandler("assets", "index.html") // if not found goes to index.html
//
func StaticHandler(assetsPath string, catch interface{}) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := "" // FilePath

		// should not be needed
		server := r.Context().Value(http.ServerContextKey).(*http.Server)
		// this is Like solving handler twice
		mux, ok := server.Handler.(*http.ServeMux)
		if ok { //
			_, handlerPath := mux.Handler(r)
			urlPath = strings.TrimPrefix(r.URL.Path, handlerPath)
		}

		sPath := path.Join(assetsPath, urlPath)

		fstat, err := os.Stat(sPath)
		if err != nil || fstat.IsDir() {
			switch t := catch.(type) {
			case http.HandlerFunc:
				t(w, r) // catchHandler
			case string:
				http.ServeFile(w, r, path.Join(assetsPath, t))
			}

			return
		}
		http.ServeFile(w, r, sPath)
	})
}

type catchHelper struct {
	http.ResponseWriter
	statusCode int
}

func (c *catchHelper) WriteHeader(code int) {
	c.statusCode = code
}

// CatchAllHandler will execute catch handler if error >= 400
// Might not work if handler uses Write on 404
func CatchAllHandler(next http.HandlerFunc, catch http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &catchHelper{w, 200}
		next(c, r)
		if c.statusCode >= 400 {
			catch(w, r)
		}
	}
}
