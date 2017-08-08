package main

import (
	"log"
	"net/http"
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

	chain := webu.NewChain(chainOne, chainTwo, chainThree)
	mux.HandleFunc("/", chain.Build(lastHandler))

	// chaining chains

	nchain := webu.NewChain(chainOne, chain.Build)
	mux.HandleFunc("/sub", nchain.Build(lastHandler))

	http.ListenAndServe(":8001", mux)

}
