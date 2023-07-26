package gee

import (
	"log"
	"net/http"
)

type Engine struct {
	*router
}

func NewEngine() *Engine {
	return &Engine{newRouter()}
}

func (m Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	context := NewContext(request, writer)
	m.handle(context)
}

func (m *Engine) Get(path string, handlerFunc handler) {
	m.addRoute("GET", path, handlerFunc)
}

func (m *Engine) Post(pattern string, handlerFunc handler) {
	m.addRoute("POST", pattern, handlerFunc)
}

func (m *Engine) Start(addr string) {
	err := http.ListenAndServe(addr, m)
	if err != nil {
		log.Fatal("start engine fail")
	}
}
