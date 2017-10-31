package webu

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gohxs/webu/chain"
)

/////////////////////////
// Log handler
/////

// LogHelper struct to handle write logs
type LogHelper struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader hijack write header to track httpStatus
func (l *LogHelper) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}

// Hijack hihack wrapper for hijacker users (websocket?)
func (l *LogHelper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := l.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}

// LogHandler returns an handler that logs output using default logger
func LogHandler(log *log.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := &LogHelper{w, 200}
		if next != nil {
			next.ServeHTTP(l, r)
		}
		raddr := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("(%s) %s %s - [%d %s]", raddr, r.Method, r.URL.Path, l.statusCode, http.StatusText(l.statusCode))
	}
}

//ChainLogger middleware for chainer
func ChainLogger(log *log.Logger) chain.Func {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return LogHandler(log, next.ServeHTTP)
	}
}
