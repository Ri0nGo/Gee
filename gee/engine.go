package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}

// Group 路由分组处理
func (rp *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rp.engine
	newGroup := &RouterGroup{
		prefix: rp.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// addRouter 将handler传递给router处理
func (rp *RouterGroup) addRouter(method HttpMethod, comp string, handler HandlerFunc) {
	pattern := rp.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	rp.engine.router.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (rp *RouterGroup) GET(pattern string, handler HandlerFunc) {
	rp.addRouter(GETMethod, pattern, handler)
}

// POST defines the method to add POST request
func (rp *RouterGroup) POST(pattern string, handler HandlerFunc) {
	rp.addRouter(POSTMethod, pattern, handler)
}

// Use 添加中间件
func (rp *RouterGroup) Use(middlewares ...HandlerFunc) {
	rp.middlewares = append(rp.middlewares, middlewares...)
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检测请求的url中是否属于含有定义了中间件
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.Contains(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, r)
	c.handlers = middlewares
	e.router.handle(c)
}
