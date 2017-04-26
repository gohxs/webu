package controller

import (
	"net/http"

	"dev.hexasoftware.com/stdio/webu"
)

type Home struct {
	Index webu.Action `webu:"index;GET,POST"`
}

func NewHome() *Home {
	c := &Home{}
	c.Index = func(w http.ResponseWriter, r *http.Request) {
		// Definition
		w.Write([]byte("Home\r\n"))
	}
	return c
}
