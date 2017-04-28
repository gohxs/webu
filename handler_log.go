package webu

import (
	"log"
	"net/http"
	"os"
)

/////////////////////////
// Log handler
/////

type logHelper struct {
	http.ResponseWriter
	statusCode int
}

func (l *logHelper) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}

func LogHandler(name string, next http.Handler) http.Handler {
	llog := log.New(os.Stderr, "["+name+"]: ", 0)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := &logHelper{w, 200}
		if next != nil {
			next.ServeHTTP(l, r)
		}
		llog.Printf("%s (%d) - %s", r.Method, l.statusCode, r.URL.Path)
	})
}

//Logger middleware for logging handlerFunc
func ChainLogger(name string) ChainFunc {
	return func(next http.Handler) http.Handler {
		return LogHandler(name, next)
	}
}
