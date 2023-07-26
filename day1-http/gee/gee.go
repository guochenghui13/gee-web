package gee

import (
	"fmt"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	routers map[string]handler
}

func NewEngine() *Engine {
	return &Engine{routers: make(map[string]handler)}
}

func (m Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	// GET, POST, PUT
	// path + get
	method := request.Method
	key := method + "-" + path

	h, ok := m.routers[key]
	if !ok {
		fmt.Fprintf(writer, "404 NOT FOUND")
		return
	}

	h(writer, request)
}

func (m *Engine) addRounte(method string, pattern string, handlerFun handler) {
	key := method + "-" + pattern
	log.Printf("addRounte %s-%s", method, pattern)
	m.routers[key] = handlerFun
}

func (m *Engine) Get(path string, handlerFunc handler) {
	m.addRounte("GET", path, handlerFunc)
}

func (m *Engine) Start(addr string) {
	err := http.ListenAndServe(addr, m)
	if err != nil {
		log.Fatal("start engine fail")
	}
}
