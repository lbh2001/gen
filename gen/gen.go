package gen

import (
	"net/http"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description: "gen is a framework analogous to gin"
 */

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

//new Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

//向路由表中添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

//GET方式添加路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST方式添加路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//实现ServeHTTP接口
//接管HTTP请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

//run
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
