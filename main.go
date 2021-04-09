package main

import (
	"fmt"
	"gen/gen"
	"net/http"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */

func main() {

	r := gen.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.PATH=%q", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]:%q\n", k, v)
		}
	})
	r.Run(":8080")

}
