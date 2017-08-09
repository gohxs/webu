package webu

import (
	"log"
	"net/http"

	"github.com/gohxs/webu/chain"
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

// LogHandler returns an handler that logs output using default logger
func LogHandler(name string, next http.HandlerFunc) http.HandlerFunc {
	//llog := log.New(os.Stderr, "["+name+"]: ", 0)
	return func(w http.ResponseWriter, r *http.Request) {
		l := &logHelper{w, 200}
		if next != nil {
			next.ServeHTTP(l, r)
		}
		log.Printf("[%s] %s %s - [%d %s]", name, r.Method, r.URL.Path, l.statusCode, http.StatusText(l.statusCode))
	}
}

//ChainLogger middleware for chainer
func ChainLogger(name string) chain.Func {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return LogHandler(name, next.ServeHTTP)
	}
}
