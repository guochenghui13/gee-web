package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter

	Path   string
	Method string
	Params map[string]string

	StatusCode int

	// middleWare
	index    int
	handlers []Handler
}

func NewContext(req *http.Request, resp http.ResponseWriter) *Context {
	return &Context{
		Req:        req,
		Resp:       resp,
		Path:       req.URL.Path,
		Method:     req.Method,
		StatusCode: -1,
		index:      -1,
		handlers:   make([]Handler, 0),
	}
}

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}

}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Resp.WriteHeader(c.StatusCode)
}

func (c *Context) SetHeader(key string, value string) {
	c.Resp.Header().Set(key, value)
}

func (c *Context) String(code int, formant string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Resp.Write([]byte(fmt.Sprintf(formant, value...)))
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Resp)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(c.Resp, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Resp.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", " text/html")
	c.Status(code)
	c.Resp.Write([]byte(html))
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.Json(code, H{"message": err})
}
