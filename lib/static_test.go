package lib_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steabert/ko/lib"
)

type CallRouter struct {
	Called *bool
}

func (s CallRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*s.Called = true
}

func TestNoneExistent(t *testing.T) {
	middleware := lib.NewStaticMiddleware(".")

	ts := httptest.NewServer(middleware(nil))
	rsp, err := http.Get(ts.URL + "/nonexistent")
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 404 {
		log.Fatal("expected file not found")
	}
}

func TestExistent(t *testing.T) {
	middleware := lib.NewStaticMiddleware(".")

	ts := httptest.NewServer(middleware(nil))
	rsp, err := http.Get(ts.URL + "/static.go")
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 200 {
		log.Fatal("expected file to be found")
	}
}

func TestIndexExist(t *testing.T) {
	middleware := lib.NewStaticMiddleware("../testdir")

	ts := httptest.NewServer(middleware(nil))
	rsp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 200 {
		log.Fatal("expected index.html to be found")
	}
}

func TestIndexNoneExist(t *testing.T) {
	middleware := lib.NewStaticMiddleware(".")

	ts := httptest.NewServer(middleware(nil))
	rsp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 404 {
		log.Fatal("expected index.html not to be found")
	}
}
