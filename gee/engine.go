package gee

import (
	"fmt"
	"net/http"
	"strconv"
)

type Engine struct {
	port   int
	routes map[string]http.HandlerFunc
}

func NewEngine(port int) *Engine {
	return &Engine{port: port, routes: make(map[string]http.HandlerFunc)}
}

func (e *Engine) AddRoutes(method string, url string, f http.HandlerFunc) {
	s := method + "-" + url
	e.routes[s] = f
}

func (e *Engine) AddGetRoutes(url string, f http.HandlerFunc) {
	e.AddRoutes("GET", url, f)
}

func (e *Engine) AddPostRoutes(url string, f http.HandlerFunc) {
	e.AddRoutes("POST", url, f)
}

func (e *Engine) AddDeleteRoutes(url string, f http.HandlerFunc) {
	e.AddRoutes("DELETE", url, f)
}

func (e *Engine) AddPutRoutes(url string, f http.HandlerFunc) {
	e.AddRoutes("PUT", url, f)
}

func (e *Engine) Serve() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := r.Method + "-" + r.URL.Path
		if f, ok := e.routes[key]; ok {
			f(w, r)
			return
		} else {
			fmt.Fprint(w, "404 NOT FOUND!")
		}
	})

	err := http.ListenAndServe(":"+strconv.Itoa(e.port), nil)
	if err != nil {
		panic(err)
	}
}
