package webu

import (
	"log"
	"net/http"
)

type Muxer interface {
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request))
}
type SubMux struct {
	Parent  Muxer
	Pattern string
	Chain   *ChainBuilder
}

func (s *SubMux) Handle(pattern string, handler http.Handler) {
	log.Println("Registering entry:", pattern)
	spath := s.Pattern

	if pattern[0] != '/' {
		pattern = "/" + pattern
	}
	spath += pattern
	/*if pattern != "" {
	}*/
	log.Println("Sub route handle", spath, " Sending to parent")
	if s.Chain != nil {
		handler = s.Chain.Build(handler)
	}

	s.Parent.Handle(spath, handler)

}
func (s *SubMux) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	s.Handle(pattern, http.HandlerFunc(handler))
}
