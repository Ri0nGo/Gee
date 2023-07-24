package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Resp http.ResponseWriter
	Req  *http.Request

	// request info
	Method HttpMethod
	Path   string
	Params map[string]string

	// response info
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Resp:   w,
		Req:    r,
		Method: HttpMethod(r.Method),
		Path:   r.URL.Path,
	}
}

// === requests 事件封装 === //
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
}

// === 内部函数封装 === //
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Resp.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Resp.Header().Set(key, value)
}

// === 提供方法，方便使用 === //

func (c *Context) JSON(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Resp)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Resp, err.Error(), 500)
	}
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Resp.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.Resp.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.Resp.Write([]byte(html))
}
