package main

import (
	"fmt"
	"github.com/lbh2001/gen/gen"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */

func main() {

	r := gen.New()

	r.Use(gen.Logger())

	g1 := r.Group("/g1")
	{
		g1.GET("/t1", func(c *gen.Context) {
			fmt.Println("执行路由方法~")
		})
	}

	g2 := r.Group("/g2")
	{
		g2.Use(gen.OnlyForGroup2())
		g2.GET("/t2", func(c *gen.Context) {
			fmt.Println("/g2/t2")
		})
		g2.GET("/z2", func(c *gen.Context) {
			fmt.Println("/g2/z2")
		})
	}

	r.Run(":9999")
	//
	//str1 := "/1/2/3"
	//str2 := "/1"
	//
	//temp1 := strings.Split(str1,"/")
	//temp2 := strings.Split(str2,"/")
	//
	//for _,v := range temp1{
	//	println(v)
	//}
	//println("-------------")
	//for _,v := range temp2{
	//	println(v)
	//}

	//r := gen.New()
	//r.GET("/login", func(c *gen.Context) {
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	c.JSON(http.StatusOK, gen.H{
	//		"password": password,
	//		"username": username,
	//		"status":   "success",
	//	})
	//})
	//
	//r.GET("/html", func(c *gen.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gen!<h1>")
	//})
	//
	//r.Run(":8998")

}
