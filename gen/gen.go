package gen

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description: gen is a simplified web framework analogous to gin
 */

// 定义 HandlerFunc
type HandlerFunc func(c *Context)

// 定义结构体
type (
	// Engine struct
	Engine struct {
		*RouterGroup
		router        *router
		groups        []*RouterGroup
		htmlTemplates *template.Template
		funcMap       template.FuncMap
	}

	// RouterGroup struct
	RouterGroup struct {
		prefix      string        //前缀
		middlewares []HandlerFunc //中间件
		engine      *Engine       //所属Engine
	}
)

// new Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// new Engine by a Default way:
// use Logger() and Recovery()
func Default() *Engine {
	engine := New()
	// 默认使用Logger和Recovery中间件
	engine.Use(Logger(), Recovery())
	return engine
}

// func Group is a initiation of a RouterGroup
// 初始化一个 RouterGroup
func (groups *RouterGroup) Group(prefix string) *RouterGroup {
	engine := groups.engine
	newGroup := &RouterGroup{
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

// 向路由表中添加路由
func (groups *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = groups.prefix + pattern
	groups.engine.router.addRoute(method, pattern, handler)
}

// GET方式添加路由
func (groups *RouterGroup) GET(pattern string, handler HandlerFunc) {
	if !(len(pattern) > 0 && pattern[0] == '/') {
		panic("路径必须以 '/'开头！")
	}
	groups.addRoute("GET", pattern, handler)
}

// POST方式添加路由
func (groups *RouterGroup) POST(pattern string, handler HandlerFunc) {
	if !(len(pattern) > 0 && pattern[0] == '/') {
		panic("路径必须以 '/'开头！")
	}
	groups.addRoute("POST", pattern, handler)
}

// 添加中间件
func (groups *RouterGroup) Use(middlewares ...HandlerFunc) {
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
	// 实例化Context
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

// Run it
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

/***************************************/
/************ HTML Template ************/
/***************************************/

// 解析对静态资源请求的路径
// 并映射为文件真实路径
// 然后返回一个对该路径请求的处理函数
func (groups *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 由相对路径得到绝对路径
	absolutePath := path.Join(groups.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// 检查文件是否存在
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// 处理对静态资源访问的请求
func (groups *RouterGroup) Static(relativePath string, root string) {
	handler := groups.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// 注册handler
	groups.GET(urlPattern, handler)
}

// 设置 engine.funcMap
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

// 加载HTML
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
