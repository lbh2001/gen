package gen

import (
	"fmt"
	"net/http"
	"strings"
)

/**
 * @Author: lbh
 * @Date: 2021/4/10
 * @Description: routers manage HandleFunc of patterns
 */

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// new router
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 将路径解析保存到切片中
func parsePattern(pattern string) []string {
	if pattern[0] != '/' {
		panic("路径必须以 '/' 开头！")
	}
	tempParts := strings.Split(pattern, "/")[1:]
	parts := make([]string, 0)

	for index, part := range tempParts {

		// 尾随斜杠重定向
		// 其将形如 /user/name/ 的路径
		// 重定向到 /user/name
		if len(part) == 0 {
			if index == len(tempParts)-1 {
				break
			} else {
				panic("路径解析错误：子路径不能为空！")
			}
		}

		parts = append(parts, part)
		if part[0] == '*' && index != len(tempParts)-1 {
			panic("在 '*'子路径后不能有其他路径！ ")
		}
	}
	return parts
}

// 向路由表中添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	key := method + "-" + pattern

	// 特殊情况: 如果完整路径是 "/"
	if pattern == "/" {
		r.roots[method].insertChild(pattern, []string{"/"}, 0)
		r.handlers[key] = handler
		fmt.Println("Registered Route: " + key)
		return
	}
	parts := parsePattern(pattern)
	// 添加结点(储存路由)
	r.roots[method].insertChild(pattern, parts, 0)

	// 储存路由对应的方法
	fmt.Println("Registered Route: " + key)
	r.handlers[key] = handler
}

// 根据方法和路径获取路由、参数
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	if path == "/" || path == "" {
		return r.roots[method].search([]string{"/"}, 0), nil
	}
	searchParts := parsePattern(path)
	params := make(map[string]string)

	// 查看router中是否有该方法(root node)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// 根据parts搜索匹配root的子结点
	resultNode := root.search(searchParts, 0)

	if resultNode != nil {
		// searchParts中每个part是精确的
		// 即pattern中的参数是被赋值了的
		// 而parts中则可能是模糊的(可能带有"/:param"等)
		// 因此二者不一致
		// 例如: addRoute时的pattern: /login/:name/home
		// 请求: host:port/login/lbh/home
		// 则: searchParts: ["login","lbh","home"]
		// 而: parts: ["login",":name","home"]
		// 需要解析addRoute时结点中的pattern
		// 其原本的参数名作为params的key
		parts := parsePattern(resultNode.pattern)

		// 储存参数
		for index, part := range parts {
			if part[0] == ':' {
				params[part] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				// 若part形如: "*someFilePath" (length > 1)
				// 说明路径为: /.../*someFilePath
				// 那么从这个part开始一直到最后一个part
				// 以"/"作为分隔符拼接成string(与parsePattern相反)
				// 作为params的key
				// 这一步做完后参数就储存完毕
				// 所以直接break
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}

		return resultNode, params

	}
	return nil, nil
}

// 获取路由并执行对应方法
func (r *router) handle(c *Context) {
	resultNode, params := r.getRoute(c.Method, c.Path)
	// 查看路由是否已被注册
	if resultNode != nil {
		c.Params = params
		// 请求中传来的真实完整路径
		realFullPath := resultNode.pattern
		key := c.Request.Method + "-" + realFullPath
		fmt.Println(key)
		// 添加到handlers中
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			ctx.JSON(http.StatusNotFound, H{
				"status": "404 Not Found",
				"method": ctx.Method,
				"path":   ctx.Path,
			})
		})
	}
	c.Next()
}

/****************************/
/******* Gen Test-Zone ******/
/*******               ******/
/*******  Test Router  ******/
/****************************/
func GetTest() *Engine {
	r := New()
	r.addRoute("GET", "/name/", catchAllHandler)
	r.addRoute("GET", "/name", catchAllHandler)
	//r.addRoute("GET", "/*", catchAllHandler)
	return r
}

func homeHandler(c *Context) {
	c.JSON(http.StatusOK, H{
		"status": "ok",
		"name":   c.Params,
	})
}

func nameHandler(c *Context) {
	c.JSON(http.StatusOK, H{
		"status": "ok",
		"page":   c.Request.URL.Path[1:] + "'s homePage",
	})
}

func catchAllHandler(c *Context) {
	c.JSON(http.StatusOK, H{
		"status": "ok",
		"name":   "this is a catchAll path handler",
	})
}
