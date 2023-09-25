package main

import (
	"gee/gee"
)

func main() {
	engine := gee.NewEngine(8080)
	engine.GET("/hello", func(c *gee.Context) {
		c.JSON(200, "hello")
	})
	engine.Serve()
}
