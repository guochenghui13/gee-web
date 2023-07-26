package main

import (
	"gee-web/day2-context/gee"
	"net/http"
)

func main() {

	engine := gee.NewEngine()

	engine.Get("/Get", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "this is server %s", "asdasdf")
	})

	engine.Post("/login", func(ctx *gee.Context) {
		ctx.Json(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"paswd":    ctx.PostForm("passwd"),
		})
	})

	engine.Start(":8080")
}
