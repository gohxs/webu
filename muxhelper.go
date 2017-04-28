package webu

import (
	"log"
	"net/http"
	"strings"
)

// BaseMuxer just a base for muxers
type MuxHandler interface {
	Handle(pattern string, handler http.Handler)
}

// Muxer webu flavoured muxer
type Muxer interface {
	MuxHandler
	Pattern(...string) string
	Group(pattern string) Muxer
	//HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request))
}

// NewMuxHelper Create a New helper mux
func NewMuxHelper(mux *http.ServeMux, pattern string, chain *ChainBuilder) Muxer {
	return &MuxBase{
		Parent:  mux,
		Chain:   chain,
		pattern: pattern,
	}
}

// Base muxer consists in a pattern and Chain
type MuxBase struct {
	Parent  MuxHandler
	Chain   *ChainBuilder
	pattern string
}

func (m *MuxBase) Handle(pattern string, handler http.Handler) {
	pattern = strings.TrimLeft(pattern, "/")
	log.Println("Solving pattern for ", pattern)
	var spath string
	if m.pattern == "/" {
		spath = "/" + pattern
	} else {
		spath = m.pattern + "/" + pattern
	}

	// Apply chain for this handler
	if m.Chain != nil {
		handler = m.Chain.Build(handler)
	}
	log.Println("Registering handle:", spath)
	//Root will handle
	m.Parent.Handle(spath, handler)

}
func (m *MuxBase) Group(pattern string) Muxer {
	return &MuxBase{
		Parent:  m,
		Chain:   nil,
		pattern: pattern,
	}
}

// Solve name
func (m *MuxBase) Pattern(sub ...string) string {

	subPattern := strings.TrimLeft(strings.Join(sub, "/"), "/")
	switch parent := m.Parent.(type) {
	case Muxer:
		log.Printf("[%s:Muxer] we add '%s'", m.pattern, subPattern)
		if subPattern == "" {
			return parent.Pattern(m.pattern) // Parent with us
		}

		if subPattern[0] == '/' {
			log.Printf("[%s:Muxer] Removing first slash", m.pattern)
			subPattern = subPattern[1:]
		}
		return parent.Pattern(m.pattern, subPattern) // With us and subPattern
	case MuxHandler:
		log.Printf("[%s:MuxHandler] Base Handler, sub: %s", m.pattern, subPattern)
		if m.pattern == "/" {
			return subPattern
		}
		return m.pattern + "/" + subPattern // us hand sub too
	}
	return ""
}

/////////////////////////////////
// Mux helper
///////////////////
/*
// MuxHelper main muxer struct
type MuxHelper struct {
	Mux     BaseMuxer // HttpMuxer
	pattern string    // Base pattern
	Chain   *ChainBuilder
}

//Handle handle http implementation
func (m *MuxHelper) Handle(pattern string, handler http.Handler) {

	spath := m.Pattern(pattern)
	log.Println("Sub route handle", spath, " Sending to parent")
	if m.Chain != nil {
		handler = m.Chain.Build(handler)
	}
	m.Mux.Handle(spath, handler)
}

//Pattern implementation
func (m *MuxHelper) Pattern(sub ...string) string {
	if len(sub) > 0 {
		return m.pattern + "/" + strings.Join(sub, "/")
	}
	return m.pattern
}

func (m *MuxHelper) Group(pattern string) *MuxGroup {
	return &MuxGroup{m, pattern, nil}
}

///////////////////////////////////
// MuxGroup
/////////////////

//MuxGroup helper to create group
type MuxGroup struct {
	Parent  Muxer
	pattern string
	Chain   *ChainBuilder
}

//Handle Implementation counting Parent muxers
func (mg *MuxGroup) Handle(pattern string, handler http.Handler) {

	spath := mg.Pattern(pattern)
	log.Println("Sub route Handler", spath, " Senting to parent")
	// Apply chain for this handler
	if mg.Chain != nil {
		handler = mg.Chain.Build(handler)
	}
	mg.Parent.Handle(spath, handler)
}

// Pattern retrieves full pattern for this group
func (mg *MuxGroup) Pattern(sub ...string) string {
	return mg.Parent.Pattern(mg.pattern, strings.Join(sub, "/"))
}

// Group sub muxer
func (mg *MuxGroup) Group(pattern string) *MuxGroup {
	return &MuxGroup{mg, pattern, nil}
}

// Built pattern
/*func (m *MuxHelper) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	m.Handle(pattern, http.HandlerFunc(handler))
}*/
