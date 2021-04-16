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

// Engine struct
type Engine struct {
	*RouteGroup
	router *router
	groups []*RouteGroup
}

// RouteGroup struct
type RouteGroup struct {
	prefix     string        //前缀
	middleware []HandlerFunc //中间件
	engine     *Engine       //所属Engine
}

//new Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.groups = []*RouteGroup{}
	engine.RouteGroup = &RouteGroup{engine: engine}
	return engine
}

// Group is a initiation of a RouteGroup
func (groups *RouteGroup) Group(prefix string) *RouteGroup {
	engine := groups.engine
	newGroup := &RouteGroup{
		prefix: groups.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//向路由表中添加路由
func (groups *RouteGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = groups.prefix + pattern
	groups.engine.router.addRoute(method, pattern, handler)
}

//GET方式添加路由
func (groups *RouteGroup) GET(pattern string, handler HandlerFunc) {
	groups.addRoute("GET", pattern, handler)
}

//POST方式添加路由
func (groups *RouteGroup) POST(pattern string, handler HandlerFunc) {
	groups.addRoute("POST", pattern, handler)
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
