package gee

import (
	"fmt"
	"net/http"
	"strconv"
)

type HandlerFunc func(*Context)

type Engine struct {
	port   int
	router *router
}

func NewEngine(port int) *Engine {
	return &Engine{port: port, router: newRouter()}
}

func (e *Engine) GET(url string, f HandlerFunc) {
	e.router.addRoute("GET", url, f)
}

func (e *Engine) POST(url string, f HandlerFunc) {
	e.router.addRoute("POST", url, f)
}

func (e *Engine) DELETE(url string, f HandlerFunc) {
	e.router.addRoute("DELETE", url, f)
}

func (e *Engine) PUT(url string, f HandlerFunc) {
	e.router.addRoute("PUT", url, f)
}

func (e *Engine) Serve() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := r.Method + "-" + r.URL.Path
		if f, ok := e.router.handlers[key]; ok {
			c := &Context{
				Writer: w,
				Req:    r,
				Path:   r.URL.Path,
				Method: r.Method,
			}
			f(c)
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
