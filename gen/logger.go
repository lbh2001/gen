package gen

import (
	"fmt"
	"time"
)

/**
 * @Author: lbh
 * @Date: 2021/4/17
 * @Description: Middleware Logger
 */

// Logger 是第一个运行的中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		// 程序起始时间
		t := time.Now()
		// 运行下一个中间件
		c.Next()
		// 打印运行状态、请求URI、程序运行时间
		fmt.Printf("[%d] %s in %v\n", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
