package main

import (
	"fmt"
	"gee/gee"
	"net/http"
)

func main() {
	engine := gee.NewEngine(8080)
	engine.AddGetRoutes("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world"+r.URL.Path)
	})
	engine.Serve()
}
