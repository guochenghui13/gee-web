package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type (
	Engine struct {
		*router
		*RouterGroup
		groups []*RouterGroup
	}

	RouterGroup struct {
		prefix      string
		middleWares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
)

func NewEngine() *Engine {
	engine := &Engine{
		router: newRouter(),
	}

	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}

	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (m Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	context := NewContext(request, writer)
	m.handle(context)
}

func (group *RouterGroup) addRoute(method string, comp string, handler handler) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (r *RouterGroup) GET(pattern string, handlerFunc handler) {
	r.addRoute("GET", pattern, handlerFunc)
}

func (r *RouterGroup) POST(pattern string, handlerFunc handler) {
	r.addRoute("POST", pattern, handlerFunc)
}

func (m *Engine) Start(addr string) {
	err := http.ListenAndServe(addr, m)
	if err != nil {
		log.Fatal("start engine fail")
	}
}
