package webu

import (
	"net/http"
	"os"
	"path"
)

// StaticHandler serves file or execute handler if file not found
func StaticHandler(assetsPath string, catchHandler http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fpath := r.URL.String()
		sPath := path.Join(assetsPath, fpath)

		fstat, err := os.Stat(sPath)
		if err != nil || fstat.IsDir() {
			catchHandler(w, r) // catchHandler
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
