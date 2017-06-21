package main

import (
	"log"
	"net/http"

	"dev.hexasoftware.com/hxs/webu"
	"dev.hexasoftware.com/hxs/webu/chain"
)

func main() {
	mux := http.NewServeMux()

	c := chain.New(webu.ChainLogger("main"))

	mux.HandleFunc("/", c.Build(webu.CatchAllHandler(func(w http.ResponseWriter, r *http.Request) {
		param := webu.Param(r)
		log.Println("Param is:", param)
		if param[0] == "hello" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Catching all")
	})))
	log.Println("Listening at :8080")
	http.ListenAndServe(":8080", mux)
}
