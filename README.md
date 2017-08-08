My web utils pkg for golang
===========================
Not a web framework but a set of helpers related for golang http

Handlers, Chainer

#### Handlers:
* Log handler
* Method handler - handlerFunc that warps POST,GET
* Special - a experimental handlerFunc to reflect methods
* Static handler - catchAll serves files, but if file does not exists execute, or return a default file

LogHandler:
```go
func main () {
	mux := http.NewMuxHelper()

	mux.HandleFunc("/",webu.LogHandler("logname", func( w http.ResponseWriter, r *http.Request ){

  }))

	http.ListenAndServer(":8080",mux)
}

```

#### TODO:
- [ ] Add tests..
- [ ] review and explain each handler/chain in documentation

