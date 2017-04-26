Web tools to easyup life
=========================


Auto controller routing
1. Focus on controller and actions, then if necessary middle wares



example code

```go
mux := http.NewServeMux()

mux.Handle(webu.CreateManager("/api/"))
```

## Checking syntax
```go
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
```
Is no diferent than
```go
type Home struct{}

func (c *Home) func( w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Home\r\n"))
}
...

func (a *App) InitRoutes() {
	h :=Home{}
	a.Router().Methods("GET","POST").Path("index").HandlerFunc(&h)	
}

	
```

