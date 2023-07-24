package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	roots    map[HttpMethod]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[HttpMethod]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRouter(method HttpMethod, pattern string, handlerFunc HandlerFunc) {
	parts := parsePattern(pattern)
	key := fmt.Sprintf("%s_%s", method, pattern)
	// 若roots中不存在node，则创建
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handlerFunc
}

func (r *Router) getRouter(method HttpMethod, path string) (*node, map[string]string) {
	// urlPath 表示用户请求的url地址
	urlPath := parsePattern(path)
	params := make(map[string]string)
	rootNode, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	node := rootNode.search(urlPath, 0)
	if node != nil {
		// node.pattern 表示的是router中的pattern 部分，即开发人员定义的api接口
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			// 若api接口中存在“:“ 或 存在 ”*“，则对api中的有名分组进行赋值，有名分组就是/api/user/:user_id 的 user_id
			if part[0] == ':' {
				params[part[1:]] = urlPath[index]
			} else if part[0] == '*' && len(parts) > 1 {
				params[part[1:]] = strings.Join(urlPath[index:], "/")
				break
			}
		}
		return node, params
	}
	return nil, nil
}

// handle http最终会调用到handle方法
func (r *Router) handle(c *Context) {
	node, params := r.getRouter(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := string(c.Method) + "_" + c.Path
		if f, ok := r.handlers[key]; ok {
			f(c)
		}
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// parsePattern 解析请求的path
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		parts = append(parts, item)
		if item[0] == '*' {
			break
		}
	}
	return parts
}
