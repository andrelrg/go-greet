package base

import (
	"log"
	"net/http"
)

type handlerFunc = func(http.ResponseWriter, *http.Request)

type Base struct {
	Handler handlerFunc
}

func (b Base) Serve(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	b.Handler(w, r)
}
