package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (e *Engine) addRouter(method HttpMethod, pattern string, handlerFunc HandlerFunc) {
	e.router.addRouter(method, pattern, handlerFunc)

}

func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRouter(GETMethod, pattern, handlerFunc)
}

func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRouter(POSTMethod, pattern, handlerFunc)
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	e.router.handle(c)
}
