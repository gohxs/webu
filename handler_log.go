package webu

import (
	"bufio"
	"log"
	"net"
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

func (l *logHelper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := l.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}

// LogHandler returns an handler that logs output using default logger
func LogHandler(log *log.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := &logHelper{w, 200}
		if next != nil {
			next.ServeHTTP(l, r)
		}
		log.Printf("%s %s - [%d %s]", r.Method, r.URL.Path, l.statusCode, http.StatusText(l.statusCode))
	}
}

//ChainLogger middleware for chainer
func ChainLogger(log *log.Logger) chain.Func {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return LogHandler(log, next.ServeHTTP)
	}
}
