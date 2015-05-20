// A [middleware](https://github.com/rkgo/web) wrapper for the high performance
// HTTP request router [httprouter](https://github.com/julienschmidt/httprouter)
//
//  routes := router.New()
//
//  routes.GET("/", auth.Index)
//  routes.GET("/logout", auth.Logout)
//  routes.POST("/login", auth.Login)
//
//  app.Use(routes.Middleware())
//
package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rkgo/web"
	"golang.org/x/net/context"
)

type key int

const paramsKey key = 0

// Handler is a function that can be registered to a route to handle
// HTTP requests
type Handler func(web.Context)

// Router can be used to dispatch requests to different handler functions
// via configurable routes
type Router struct {
	router httprouter.Router
}

// New returns a new initialized Router.
func New() *Router {
	return &Router{
		router: httprouter.Router{
			RedirectTrailingSlash: true,
			RedirectFixedPath:     true,
			NotFound: func(rw http.ResponseWriter, _ *http.Request) {
				// do nothing
			},
		},
	}
}

// GET is a shortcut for router.Handle("GET", path, handle)
func (r *Router) GET(path string, handler Handler) {
	r.Handle("GET", path, handler)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func (r *Router) HEAD(path string, handler Handler) {
	r.Handle("HEAD", path, handler)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func (r *Router) OPTIONS(path string, handler Handler) {
	r.Handle("OPTIONS", path, handler)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func (r *Router) POST(path string, handler Handler) {
	r.Handle("POST", path, handler)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func (r *Router) PUT(path string, handler Handler) {
	r.Handle("PUT", path, handler)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func (r *Router) PATCH(path string, handler Handler) {
	r.Handle("PATCH", path, handler)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func (r *Router) DELETE(path string, handler Handler) {
	r.Handle("DELETE", path, handler)
}

// Handle registers a new request handle with the given path and method.
func (r *Router) Handle(method, path string, handler Handler) {
	r.router.Handle(method, path, func(rw http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		ctx, ok := rw.(web.Context)
		if !ok {
			panic(fmt.Errorf("invalid context"))
		}

		handler(ctx.WithValue(paramsKey, params))
	})
}

// Param reads the parameter for the given name from the provided context. The
// result will be an empty string, if the parameter does not exist.
func Param(ctx context.Context, name string) string {
	params, ok := ctx.Value(paramsKey).(httprouter.Params)
	if !ok {
		return ""
	}

	return params.ByName(name)
}

// Middleware returns a [rkgo/web](https://github.com/rkgo/web) compatible
// middleware
func (r *Router) Middleware() web.Middleware {
	return func(ctx web.Context, next web.Next) {
		r.router.ServeHTTP(ctx, ctx.Req())

		if ctx.Status() == 0 {
			next(ctx)
		}
	}
}