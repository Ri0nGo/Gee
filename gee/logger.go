package gee

import (
	"fmt"
	"log"
	"time"
)

// Logger 日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		startTime := time.Now()
		c.Next()
		useTime := time.Now().Sub(startTime).Seconds()
		log.Println(fmt.Sprintf("code: %d, path: %s, use time: %.3f",
			c.StatusCode, c.Path, useTime))
	}
}
