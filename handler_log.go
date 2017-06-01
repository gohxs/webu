package webu

import (
	"log"
	"net/http"

	"dev.hexasoftware.com/hxs/webu/chain"
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

func LogHandler(name string, next http.HandlerFunc) http.HandlerFunc {
	//llog := log.New(os.Stderr, "["+name+"]: ", 0)
	return func(w http.ResponseWriter, r *http.Request) {
		l := &logHelper{w, 200}
		if next != nil {
			next.ServeHTTP(l, r)
		}
		log.Printf("[%s] %s (%d) - %s", name, r.Method, l.statusCode, r.URL.Path)
	}
}

//Logger middleware for logging handlerFunc
func ChainLogger(name string) chain.Func {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return LogHandler(name, next.ServeHTTP)
	}
}
