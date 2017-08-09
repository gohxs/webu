package webu_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gohxs/webu"
)

func ExampleParams() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := webu.Params(r)
		fmt.Println("Param:", len(p), p)
	})
	s := httptest.NewServer(mux)

	http.Get(s.URL + "/1/10/20")
	http.Get(s.URL + "/a1/test")
	// Output:
	// Param: 3 [1 10 20]
	// Param: 2 [a1 test]
}

func ExampleParams2() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/thing/", func(w http.ResponseWriter, r *http.Request) {
		p := webu.Params(r)
		fmt.Println("Param:", len(p), p)
	})
	s := httptest.NewServer(mux)
	http.Get(s.URL + "/api/v1/thing/id")
	http.Get(s.URL + "/api/v1/thing/2")
	http.Get(s.URL + "/api/v1/thing")
	// Output:
	// Param: 1 [id]
	// Param: 1 [2]
	// Param: 0 []
}

func ExampleParamsQuery() {
	mux := http.NewServeMux()
	mux.HandleFunc("/query/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Param:", webu.Params(r))
		fmt.Printf("Query: a:%s b:%s\n", r.URL.Query()["a"], r.URL.Query()["b"])
	})
	s := httptest.NewServer(mux)
	http.Get(s.URL + "/query/test?a=1&b=2&a=2")
	http.Get(s.URL + "/query/ok")
	// Output:
	// Param: [test]
	// Query: a:[1 2] b:[2]
	// Param: [ok]
	// Query: a:[] b:[]

}
