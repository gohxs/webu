My web utils pkg for golang
===========================


Todo:
Add special case directly to router

```go
m.Handle("/test", webu.SpecialHandler(m.Pattern("/test"), SpecialFactory, "Index"))
// AS:
m.HandleSpecial("/test",Factory,"Method")
```

