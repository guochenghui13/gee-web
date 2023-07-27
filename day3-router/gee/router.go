package gee

import (
	"log"
	"net/http"
	"strings"
)

type handler func(ctx *Context)

type router struct {
	roots    map[string]*node
	handlers map[string]handler
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]handler)}
}

// we only allow one *
func parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, s := range split {
		if s != "" {
			parts = append(parts, s)
			if s[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (m *router) addRoute(method string, pattern string, handlerFun handler) {
	key := method + "-" + pattern
	log.Printf("addRoute %s-%s", method, pattern)

	_, ok := m.roots[method]
	if !ok {
		m.roots[method] = &node{}
	}
	m.roots[method].Insert(pattern, parsePattern(pattern), 0)
	m.handlers[key] = handlerFun
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.Search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
