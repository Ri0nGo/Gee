package gee

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRouter(method HttpMethod, pattern string, handlerFunc HandlerFunc) {
	key := fmt.Sprintf("%s_%s", method, pattern)
	r.handlers[key] = handlerFunc
}

func (r *Router) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if f, ok := r.handlers[key]; ok {
		f(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
