package main

import (
	"fmt"
	"gee-web/day4-group/gee"
	"net/http"
)

func main() {
	engine := gee.NewEngine()

	group := engine.Group("/admin")
	{
		group.GET("/query", func(ctx *gee.Context) {
			fmt.Println("/admin/query")
			ctx.String(http.StatusOK, "this is admin %s", "gee")
		})

		groupData := group.Group("/biz")
		{
			groupData.GET("/data", func(ctx *gee.Context) {
				ctx.String(http.StatusOK, "this is data bu")
			})
		}
	}

	engine.GET("/index", func(ctx *gee.Context) {
		fmt.Println("/index")
		ctx.String(http.StatusOK, "index %s", "asd ")
	})

	engine.Start(":8080")

}
