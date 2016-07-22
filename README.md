# router

A middleware that works well (but not exclusively) with [rkusa/web](https://github.com/rkusa/web) and provides high performance HTTP request routing using [httptreemux](http://github.com/dimfeld/httptreemux).

A [middleware](https://github.com/rkgo/web) wrapper for the high performance HTTP request router [httptreemux](http://github.com/dimfeld/httptreemux)

[![Build Status][travis]](https://travis-ci.org/rkusa/router)
[![GoDoc][godoc]](https://godoc.org/github.com/rkusa/router)

### Example

```go
routes := router.New()

routes.GET("/users/:id", func(rw http.ResponseWriter, r *http.Request) {
  id := router.Param(r, "id")
  // ...
})

app.Use(routes.Middleware())
```

## License

[MIT](LICENSE)

[travis]: https://img.shields.io/travis/rkusa/router.svg
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg
