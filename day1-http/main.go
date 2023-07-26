package main

import (
	"fmt"
	"gee-web/day1-http/gee"
	"log"
	"net/http"
)

func main() {

	engine := gee.NewEngine()
	engine.Get("/get", func(w http.ResponseWriter, r *http.Request) {

		_, err := fmt.Fprintf(w, "engine created")
		if err != nil {
			log.Fatal("write error")
		}
	})

	engine.Start(":8080")
}
