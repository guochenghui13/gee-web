package main

/*
(1) global middleware Logger
$ curl http://localhost:9999/
<h1>Hello Gee</h1>

>>> log
2019/08/17 01:37:38 [200] / in 3.14µs
*/

/*
(2) global + group middleware
$ curl http://localhost:9999/v2/hello/geektutu
{"message":"Internal Server Error"}

>>> log
2019/08/17 01:38:48 [200] /v2/hello/geektutu in 61.467µs for group v2
2019/08/17 01:38:48 [200] /v2/hello/geektutu in 281µs
*/

import (
	gee2 "gee-web/day5-middleware/gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee2.Handler {
	return func(c *gee2.Context) {
		// Start timer
		t := time.Now()

		c.Next()

		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee2.NewEngine()
	r.Use(onlyForV2()) // global midlleware
	r.GET("/", func(c *gee2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee2.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Start(":8080")
}
