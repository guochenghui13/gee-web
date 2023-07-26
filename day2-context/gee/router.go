package gee

import (
	"log"
	"net/http"
)

type handler func(ctx *Context)

type router struct {
	handlers map[string]handler
}

func newRouter() *router {
	return &router{handlers: make(map[string]handler)}
}

func (m *router) addRoute(method string, pattern string, handlerFun handler) {
	key := method + "-" + pattern
	log.Printf("addRoute %s-%s", method, pattern)
	m.handlers[key] = handlerFun
}

func (m *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	h, ok := m.handlers[key]
	if !ok {
		http.Error(c.Resp, "not found", http.StatusNotFound)
		return
	}

	h(c)
}
