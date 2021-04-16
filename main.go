package main

import "gen/gen"

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */

func main() {

	r := gen.GetTest()
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
