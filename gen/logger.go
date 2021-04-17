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

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		fmt.Printf("[%d] %s in %v\n", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
