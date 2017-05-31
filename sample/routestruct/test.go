package main

import (
	"log"
	"net/http"

	"dev.hexasoftware.com/hxs/webu"
)

func MuxRegister(mux webu.Muxer, routes webu.Routes) {

	var addRoutes func(prefix string, e *webu.Entry, children webu.Routes)

	addRoutes = func(prefix string, e *webu.Entry, children webu.Routes) {
		for _, r := range children {
			log.Println("Path:", prefix+r.Path)
			if r.Handler == nil {
				log.Println("No handler for this route")
			} else {
				handleFunc := r.Handler
				if r.Methods != nil {
					th := webu.Method{}
					for _, m := range r.Methods {
						log.Println("Registering for method:", m)
						th[m] = r.Handler
					}
					handleFunc = th.ServeHTTP
				}
				// ChainBuild handleFunc
				if r.Chain != nil {
					handleFunc = r.Chain.Build(handleFunc)
				}
				// No parent entry
				mux.Handle(prefix+r.Path, handleFunc)
			}
			// Sub Children
			if r.Children != nil {
				addRoutes(prefix+r.Path, &r, r.Children)
			}
		}
	}
	addRoutes("", nil, routes)
}

func DBGHandler(name string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("This is handler", name)
		w.Write([]byte(name))
	}
}

func main() {
	routes := webu.Routes{
		webu.Entry{
			Path:  "/api/v1",
			Chain: webu.NewChain(webu.ChainLogger("api")),
			Children: webu.Routes{
				webu.Entry{
					Path: "/user",
					Children: webu.Routes{
						webu.Entry{Path: "/login", Handler: DBGHandler("/login")},
					},
				},
			},
		},
	}

	mux := http.NewServeMux()

	MuxRegister(mux, routes)

	http.ListenAndServe(":3001", mux)

}
