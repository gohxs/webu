package main

import (
	"log"
	"net/http"
	"strings"

	"dev.hexasoftware.com/hxs/webu"
)

// The Routing definition for controller
func InitHome(m webu.Muxer) {

	h := &HomeHandler{}
	indexHandler := &webu.MethodHandler{Get: h.Index, Post: h.POSTIndex}
	m.Handle("/", indexHandler)
	m.Handle("/special/", webu.SpecialHandler("/home/special/", h.Special))
}

type HomeHandler struct{}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET GET Hello\r\n"))
}
func (h *HomeHandler) POSTIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST POST POST Hello\r\n"))
}
func (h *HomeHandler) Special(params ...string) (interface{}, error) {
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

/////////////////////////////////////////////////
//// MAIN ////
/////////////////////////
func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	mux := http.NewServeMux()
	// Sub route name

	// golang mux compatible
	chain := webu.NewChain(webu.ChainLogger("HOME"))

	// Controller somehow
	InitHome(&webu.SubMux{mux, "/home", chain})

	// Catch all
	//mux.Handle("/", webu.LogHandler(http.HandlerFunc(notFoundHandler)))

	http.ListenAndServe(":8080", mux)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found \r\n"))
}

/*






































 */
//////////////////////////////////
// Routing test
/////////////
type RouteManager interface {
	Handle(name string, handler http.HandlerFunc)
	Route(name string) RouteManager // Child routing
	Name() string
}

// Router sub router
/////////////////////////////////////////
type Router struct { //Just a sub handler
	name   string
	parent RouteManager
}

func (r *Router) Handle(entry string, handler http.HandlerFunc) {
	log.Println("Registering entry:", entry)
	name := entry
	spath := r.name

	p := strings.Split(entry, "#")
	if len(p) > 1 {
		spath = p[0] + "#/" + r.name
		name = p[1]
	}
	if name[0] != '/' {
		name = "/" + name
	}
	if name != "" {
		spath += name
	}
	log.Println("Sub route handle", spath, " Sending to parent")
	r.parent.Handle(spath, handler)
}
func (r *Router) Route(name string) RouteManager {
	return &Router{name, r}
}
func (r *Router) Name() string {
	return r.parent.Name() + "/" + r.name
}

////////////////////////////////////////
// Master router
///
type CoreRouter struct {
	entry map[string]http.HandlerFunc
}

func NewCoreRouter() *CoreRouter {
	return &CoreRouter{map[string]http.HandlerFunc{}}
}

func (c *CoreRouter) Handle(entry string, handler http.HandlerFunc) {
	log.Println("Registering entry", entry)
	name := entry
	methods := []string{}

	p := strings.Split(entry, "#")
	if len(p) > 1 {
		log.Println("Methods existent", p[1])
		methods = strings.Split(p[0], ",")
		name = p[1]
	}
	// Prefix if not prefixed
	if name[0] != '/' {
		name = "/" + name
	}

	name = strings.TrimRight(name, "/")
	if len(methods) > 0 {
		for _, v := range methods {
			entryStr := v + "#" + name
			log.Println("Registering:", entryStr)
			c.entry[entryStr] = handler
		}
		return
	}
	c.entry[name] = handler

}
func (c *CoreRouter) Name() string {
	return ""
}

func (c *CoreRouter) Solve(method string, path string) http.HandlerFunc {

	log.Println("Entries:", c.entry)
	entryStr := method + "#" + path
	log.Println("Entry is :", entryStr)
	handler, ok := c.entry[entryStr]
	if ok {
		log.Println("Route found")
		return handler
	}
	// If not found we try without method
	handler = c.entry[path]
	return handler
}
func (c *CoreRouter) Route(name string) RouteManager {
	return &Router{name, c}
}

// Method#path/params
func (c *CoreRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	vpath := strings.TrimRight(r.URL.Path, "/")
	handler := c.Solve(r.Method, vpath)
	if handler != nil {
		handler(w, r)
		return
	}
	w.WriteHeader(404)
	w.Write([]byte("Not Found"))

	//
}
