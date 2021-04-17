package gen

import (
	"fmt"
)

/**
 * @Author: lbh
 * @Date: 2021/4/17
 * @Description:
 */

func Test1() HandlerFunc {
	return func(c *Context) {
		fmt.Println("here is part1")
		c.Next()
		fmt.Println("here is part4")
	}
}

func Test2() HandlerFunc {
	return func(c *Context) {
		fmt.Println("here is part2")
		c.Next()
		fmt.Println("here is part3")
	}
}

func OnlyForGroup2() HandlerFunc {
	return func(c *Context) {
		fmt.Println("this is group2 middleware")
		c.Next()
	}
}
