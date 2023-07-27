package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/hello/geektutu"
hello geektutu, you're at /hello/geektutu

(4)
$ curl "http://localhost:9999/assets/css/geektutu.css"
{"filepath":"css/geektutu.css"}

(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"gee-web/day2-context/gee"
	gee2 "gee-web/day3-router/gee"
	"net/http"
)

func main() {
	r := gee2.NewEngine()
	r.Get("/", func(c *gee2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Get("/hello", func(c *gee2.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/hello/:name", func(c *gee2.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *gee2.Context) {
		c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Start(":9999")
}
