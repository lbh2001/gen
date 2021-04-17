package gen

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

/**
 * @Author: lbh
 * @Date: 2021/4/17
 * @Description: Middleware Recovery
 */

// 捕获panic的信息用以输出
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// Recovery 防止panic后程序就结束
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				errorMessage := fmt.Sprintf("error: %s", err)
				log.Printf("%s\n\n", trace(errorMessage))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
