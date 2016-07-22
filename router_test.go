package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rkusa/web"
)

func TestMiddleware(t *testing.T) {
	r := New()
	r.GET("/foo", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("bar"))
	})

	rec := request(t, r, "GET", "/foo")

	if rec.Code != http.StatusOK {
		t.Errorf("request failed")
	}

	if rec.Body.String() != "bar" {
		t.Errorf("unexpected body")
	}
}

func TestNotFound(t *testing.T) {
	r := New()
	r.GET("/foo", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("bar"))
	})

	rec := request(t, r, "GET", "/bar")

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got: %d", rec.Code)
	}
}

func TestGroup(t *testing.T) {
	r := New()
	a := r.Group("/a")
	{
		a.GET("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("/a"))
		})

		a.GET("/b", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("/a/b"))
		})

		a.GET("c", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("/a/c"))
		})

		d := a.Group("d")
		{
			d.GET("", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("/a/d"))
			})

			d.GET("/e", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("/a/d/e"))
			})
		}

		f := a.Group("f/")
		{
			f.GET("/g", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("/a/f/g"))
			})
		}
	}

	paths := []string{
		"/a", "/a/b", "/a/c", "/a/d", "/a/d/e", "/a/f/g",
	}

	for _, path := range paths {
		rec := request(t, r, "GET", path)
		if rec.Code != http.StatusOK || rec.Body.String() != path {
			t.Errorf(
				"path %s not working as expected: status: %d, body: %s",
				path, rec.Code, rec.Body.String(),
			)
		}
	}

	rec := request(t, r, "GET", "/a/f")
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got: %d", rec.Code)
	}
}

func request(t *testing.T, router *Router, method, path string) *httptest.ResponseRecorder {
	app := web.New()
	app.Use(router.Middleware())

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	app.ServeHTTP(rec, req)

	return rec
}
