package main

import (
	"Gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/home", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/index", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"username": c.PostForm("user"),
			"password": c.PostForm("pwd"),
		})
	})

	r.Run(":8000")
}
