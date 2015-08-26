# router

A [middleware](https://github.com/rkgo/web) wrapper for the high performance HTTP request router [httptreemux](http://github.com/dimfeld/httptreemux)

[![Build Status][drone]](https://ci.rkusa.st/github.com/rkgo/router)
[![GoDoc][godoc]](https://godoc.org/github.com/rkgo/router)

### Example

```go
routes := router.New()

routes.GET("/", auth.Index)
routes.GET("/logout", auth.Logout)
routes.POST("/login", auth.Login)

app.Use(routes.Middleware())
```

[drone]: http://ci.rkusa.st/api/badge/github.com/rkgo/router/status.svg?branch=master&style=flat-square
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square