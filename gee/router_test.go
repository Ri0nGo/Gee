package gee

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestRouter() *Router {
	r := NewRouter()
	fmt.Println("test new router: ", r)
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/`hello`/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

// TestParsePattern 测试解析url函数
func TestParsePattern(t *testing.T) {
	fmt.Println(parsePattern("/p/:name"))
	fmt.Println(parsePattern("/p/*"))
	fmt.Println(parsePattern("/p/*name/*"))
	assert.Equal(t, parsePattern("/p/:name"), []string{"p", ":name"})
	assert.Equal(t, parsePattern("/p/*"), []string{"p", "*"})
	assert.Equal(t, parsePattern("/p/*name/*"), []string{"p", "*name"})
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	fmt.Println("new router", r)
	n, ps := r.getRouter("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}
