package main

import (
	"log"
	"net/http"

	"github.com/gohxs/webu/chain"
)

func chainOne(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("One")
		next(w, r)
		log.Println("After one")
	}
}
func chainTwo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Two")
		next(w, r)
		log.Println("After two")
	}
}
func chainThree(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Three")
		next(w, r)
		log.Println("After three")
	}
}

func lastHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Final")
}

func main() {

	mux := http.NewServeMux()

	chn := chain.New(chainOne, chainTwo, chainThree)
	mux.HandleFunc("/", chn.Build(lastHandler))

	// chaining chains

	nchn := chain.New(chainOne, chn.Build)
	mux.HandleFunc("/sub", nchn.Build(lastHandler))

	http.ListenAndServe(":8001", mux)

}
