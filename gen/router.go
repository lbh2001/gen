package gen

import (
	"net/http"
	"strings"
)

/**
 * @Author: lbh
 * @Date: 2021/4/10
 * @Description:
 */

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

//new router
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//将路径解析保存到切片中
func parsePattern(pattern string) []string {
	tempParts := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, part := range tempParts {
		parts = append(parts, part)
		if part[0] == '*' {
			return parts
		}
	}
	return parts
}

//向路由表中添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	//添加结点(储存路由)
	r.roots[method].insert(pattern, parts, 0)

	//储存路由对应的方法
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//根据方法和路径获取路由、参数
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	//查看router中是否有该方法(root node)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	//根据parts搜索匹配root的子结点
	resultNode := root.search(searchParts, 0)

	if resultNode != nil {
		//searchParts中每个part是精确的
		//即pattern中的参数是被赋值了的
		//而parts中则可能是模糊的(可能带有"/:param"等)
		//因此二者不一致
		//例如: addRoute时的pattern: /login/:name/home
		//请求: host:port/login/lbh/home
		//则: searchParts: ["login","lbh","home"]
		//而: parts: ["login",":name","home"]
		//需要解析addRoute时结点中的pattern
		//其原本的参数名作为params的key
		parts := parsePattern(resultNode.pattern)

		//储存参数
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				//若part形如: "*someFilePath" (length > 1)
				//说明路径为: /.../*someFilePath/...
				//那么从这个part开始一直到最后一个part
				//以"/"作为分隔符拼接成string(与parsePattern相反)
				//作为params的key
				//这一步做完后参数就储存完毕
				//所以直接break
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}

		return resultNode, params

	}
	return nil, nil
}

//获取路由并执行对应方法
func (r *router) handle(c *Context) {
	resultNode, params := r.getRoute(c.Method, c.Path)
	//查看路由是否已被注册
	if resultNode != nil {
		c.Params = params
		key := c.Request.Method + "-" + c.Request.URL.Path
		//执行方法
		r.handlers[key](c)
	} else {
		c.JSON(http.StatusNotFound, H{
			"status": "404 Not Found",
			"method": c.Method,
			"path":   c.Path,
		})
	}

}

/****************************/
/******* Gen Test-Zone ******/
/*******               ******/
/*******  Test Router  ******/
/****************************/
func getTest() *router {
	r := newRouter()

	return r
}
