package main

import (
	"Gee/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path %s, Method: %s \n", r.Method, r.URL)
	})

	r.POST("/index", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header [%s] = %s \n", k, v)
		}
	})

	r.Run(":8000")
}
