package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRouter(method HttpMethod, pattern string, handlerFunc HandlerFunc) {
	key := fmt.Sprintf("%s_%s", method, pattern)
	e.router[key] = handlerFunc

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
	key := r.Method + "_" + r.URL.Path
	if f, ok := e.router[key]; ok {
		f(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Request path 404 NOT FOUND: %s\n", r.URL)
	}
}
