package test

import (
	"fmt"
	"github.com/lbh2001/gen/gen"
)

/**
 * @Author: lbh
 * @Date: 2021/4/17
 * @Description: Middlewares test
 */

func Tests() gen.HandlerFunc {
	return func(c *gen.Context) {
		fmt.Println("here is part1")
		c.Next()
		fmt.Println("here is part4")
	}
}

func Test() gen.HandlerFunc {
	return func(c *gen.Context) {
		fmt.Println("here is part2")
		c.Next()
		fmt.Println("here is part3")
	}
}

func OnlyForGroup2() gen.HandlerFunc {
	return func(c *gen.Context) {
		fmt.Println("this is group2 middleware")
		c.Next()
	}
}
