package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rkgo/web"
)

func TestMiddleware(t *testing.T) {
	rec := request(t, "GET", "/foo")

	if rec.Code != http.StatusOK {
		t.Errorf("request failed")
	}

	if rec.Body.String() != "bar" {
		t.Errorf("unexpected body")
	}
}

func TestNotFound(t *testing.T) {
	rec := request(t, "GET", "/bar")

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got: %d", rec.Code)
	}
}

func request(t *testing.T, method, path string) *httptest.ResponseRecorder {
	router := New()
	router.GET("/foo", func(ctx web.Context) {
		ctx.Write([]byte("bar"))
	})

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
