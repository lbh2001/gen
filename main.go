package main

import (
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

	r.GET("/login", func(c *gen.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gen.H{
			"password": password,
			"username": username,
			"status":   "success",
		})
	})

	r.GET("/html", func(c *gen.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gen!<h1>")
	})

	r.Run(":8998")

}
