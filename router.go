// Package router is a middleware that works well (but not exclusively) with
// [rkusa/web](https://github.com/rkusa/web) and provides high performance HTTP
// request routing using [httptreemux](http://github.com/dimfeld/httptreemux).
//
//  routes := router.New()
//
//  routes.GET("/users/:id", func(rw http.ResponseWriter, r *http.Request) {
//    id := router.Param(r, "id")
//    // ...
//  })
//
//  app.Use(routes.Middleware())
//
package router

import (
	"context"
	"net/http"
	pathhelper "path"

	"github.com/dimfeld/httptreemux"
)

type key int

const paramsKey key = 0

// Router can be used to dispatch requests to different handler functions
// via configurable routes
type Router struct {
	router   httptreemux.TreeMux
	basePath string
}

// New returns a new initialized Router.
func New() *Router {
	r := &Router{
		router:   *httptreemux.New(),
		basePath: "",
	}
	r.router.PathSource = httptreemux.URLPath
	r.router.NotFoundHandler = func(rw http.ResponseWriter, _ *http.Request) {
		// do nothing
	}
	return r
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (rt *Router) GET(path string, handler http.HandlerFunc) {
	rt.Handle("GET", path, handler)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func (rt *Router) HEAD(path string, handler http.HandlerFunc) {
	rt.Handle("HEAD", path, handler)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func (rt *Router) OPTIONS(path string, handler http.HandlerFunc) {
	rt.Handle("OPTIONS", path, handler)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func (rt *Router) POST(path string, handler http.HandlerFunc) {
	rt.Handle("POST", path, handler)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func (rt *Router) PUT(path string, handler http.HandlerFunc) {
	rt.Handle("PUT", path, handler)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func (rt *Router) PATCH(path string, handler http.HandlerFunc) {
	rt.Handle("PATCH", path, handler)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func (rt *Router) DELETE(path string, handler http.HandlerFunc) {
	rt.Handle("DELETE", path, handler)
}

// Handle registers a new request handle with the given path and method.
func (rt *Router) Handle(method, path string, handler http.HandlerFunc) {
	rt.router.Handle(
		method,
		pathhelper.Join(rt.basePath, path),
		func(rw http.ResponseWriter, r *http.Request, params map[string]string) {
			handler(rw, r.WithContext(context.WithValue(r.Context(), paramsKey, params)))
		},
	)
}

// Group creates a new route group with the given path prefix. All route created
// using the returned Router are prefixed accoringly.
func (rt *Router) Group(path string) *Router {
	return &Router{
		router:   rt.router,
		basePath: pathhelper.Join(rt.basePath, path),
	}
}

// Param reads the parameter for the given name from the provided context. The
// result will be an empty string, if the parameter does not exist.
func Param(r *http.Request, name string) string {
	params, ok := r.Context().Value(paramsKey).(map[string]string)
	if !ok {
		return ""
	}

	return params[name]
}

type withWritten interface {
	http.ResponseWriter
	Written() bool
}

// Middleware returns a middleware that serves all registered routes.
func (rt *Router) Middleware() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		rt.router.ServeHTTP(rw, r)

		if rw, ok := rw.(withWritten); ok && !rw.Written() {
			next(rw, r)
		}
	}
}
