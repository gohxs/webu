My web utils pkg for golang
===========================
Focusing on chainer def golang middleware


Todo:
Add special case directly to router

```go
m.Handle("/test", webu.SpecialHandler(m.Pattern("/test"), SpecialFactory, "Index"))
// AS:
m.HandleSpecial("/test",Factory,"Method")
```

Changes:
20-06-2017 Added fetchParam function


