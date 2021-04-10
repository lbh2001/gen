package gen

import (
	"net/http"
)

/**
 * @Author: lbh
 * @Date: 2021/4/10
 * @Description:
 */

type router struct {
	handlers map[string]HandlerFunc
}

//new router
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

//向路由表中添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//根据路由表转发路由
func (r *router) handle(c *Context) {
	key := c.Request.Method + "-" + c.Request.URL.Path

	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.JSON(http.StatusNotFound, H{
			"method":  c.Request.Method,
			"url":     c.Request.URL.Path,
			"message": "404 Not Found",
		})
	}
}
