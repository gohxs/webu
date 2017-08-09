package webu_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gohxs/webu"
)

func ExampleMethods() {

	mux := http.NewServeMux()
	mux.Handle("/", webu.Methods{
		"GET": func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HandleGet")
		},
		"POST": func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HandlePost")
		},
	})
	s := httptest.NewServer(mux)
	http.Get(s.URL)
	http.Post(s.URL, "", bytes.NewBuffer([]byte{1, 2}))
	// Output:
	// HandleGet
	// HandlePost
}

func ExampleMethodFunc() {
	mux := http.NewServeMux()
	methods := webu.Methods{}
	methods.Get(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HandleGet")
	})
	methods.Post(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HandlePost")
	})

	mux.Handle("/", methods)

	s := httptest.NewServer(mux)
	http.Get(s.URL)
	http.Post(s.URL, "", bytes.NewBuffer([]byte{1, 2}))
	// Output:
	// HandleGet
	// HandlePost

}
