package main

import (
	"log"
	"net/http"

	"dev.hexasoftware.com/stdio/webu"
	"dev.hexasoftware.com/stdio/webu/sample/test1/controller"
)

func main() {
	mux := http.NewServeMux()

	mgr := webu.CreateManager("api")

	mgr.AddController("home", controller.NewHome())

	mux.Handle(mgr.Handler())
	log.Println("Listening at :8080")
	http.ListenAndServe(":8080", mux)
}
