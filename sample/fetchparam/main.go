package main

import (
	"log"
	"net/http"

	"github.com/gohxs/webu"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		extra := webu.FetchParam(r)
		log.Println("Extra:", extra)
	})

	log.Println("Listening at :8081")
	http.ListenAndServe(":8081", mux)
}
