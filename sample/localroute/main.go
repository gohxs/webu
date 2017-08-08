package main

import (
	"log"
	"net/http"

	"github.com/gohxs/prettylog"
	"github.com/gohxs/webu"
)

// The Routing definition for controller
func RouteHome(m webu.Muxer) {
	log.Println("Muxer name:", m.Pattern())
	h := &HomeHandler{}
	indexHandler := webu.MethodHandler{"GET": h.Index, "POST": h.POSTIndex}
	m.Handle("/", indexHandler)
	m.Handle("", indexHandler)
}

type HomeHandler struct{}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Context:", r.Context().Value("path"))
	w.Write([]byte("GET GET Hello\r\n"))
}
func (h *HomeHandler) POSTIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST POST POST Hello\r\n"))
}

//////////////////
// Special handler, struct per request
/////
func RouteSpecial(m webu.Muxer) {
	m.Handle("/", webu.SpecialHandler(m.Pattern("/"), SpecialFactory, "Index"))
	m.Handle("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Parameter needed\r\n"))
	}))

	m.Handle("/test/", webu.SpecialHandlerFunc(m.Pattern("/test/"), func(params ...string) (interface{}, error) {
		log.Println("Test params:", params)
		return params, nil
	}))
}

type SpecialHandler struct { // Created every request
	*webu.SpecialBase
}

func SpecialFactory(base *webu.SpecialBase) webu.SpecialInterface {
	return &SpecialHandler{base}
}

func (s *SpecialHandler) Index(params ...string) (interface{}, error) {
	log.Println("What is my url??", s.R.URL.Path)
	log.Println("Params:", params)
	id := ""
	if len(params) >= 1 {
		id = params[0]
	}
	return map[string]interface{}{
		"Result": "Fine",
		"ID":     id,
	}, nil
}

func ChainAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Headers:", r.Header)
		authstr := r.Header.Get("Authorization")
		if authstr != "Basic c3RkaW86MXEydzNl" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Authorization fail, but this is a test, we will let you go trough with user:stdio pass:1q2w3e\r\n"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

/////////////////////////////////////////////////
//// MAIN ////
/////////////////////////
func main() {
	prettylog.Global()

	mux := http.NewServeMux()
	// Sub route name

	// golang mux compatible
	//chain := webu.NewChain(webu.ChainLogger("HOME"))
	muxer := webu.NewMuxHelper(mux, nil)
	//subMuxer := muxer.Group("/home").Group("/test").Group("sub")
	// Controller somehow
	RouteHome(muxer.Group("/home", webu.NewChain(webu.ChainLogger("HOME"), ChainAuth)))
	RouteSpecial(muxer.Group("/special", webu.NewChain(webu.ChainLogger("SPECIAL"))))

	// Catch all
	//mux.Handle("/", webu.LogHandler(http.HandlerFunc(notFoundHandler)))

	http.ListenAndServe(":8080", mux)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found \r\n"))
}
