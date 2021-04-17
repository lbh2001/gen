package gen

import (
	"fmt"
	"net/http"
	"strings"
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
	prefix      string        //前缀
	middlewares []HandlerFunc //中间件
	engine      *Engine       //所属Engine
}

// new Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

// func Group is a initiation of a RouteGroup
func (groups *RouteGroup) Group(prefix string) *RouteGroup {
	engine := groups.engine
	newGroup := &RouteGroup{
		prefix: groups.prefix + prefix,
		engine: engine,
	}
	// 路由组不能重复(路径唯一性)
	for _, group := range engine.groups {
		if group.prefix == newGroup.prefix {
			panic(fmt.Sprintf("路由组 \"%v\" 已存在！", group.prefix))
		}
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

// Add middleware
func (groups *RouteGroup) Use(middlewares ...HandlerFunc) {
	groups.middlewares = append(groups.middlewares, middlewares...)
}

// 实现ServeHTTP接口
// 接管所有HTTP请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	// 添加要执行的中间件
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}

// Run it
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
