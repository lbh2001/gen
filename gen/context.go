package gen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * @Author: lbh
 * @Date: 2021/4/10
 * @Description:
 */

//简化
type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Method     string            //请求方法
	Path       string            //请求路径
	Params     map[string]string //解析后的参数
	StatusCode int               //状态码
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Method:  req.Method,
		Path:    req.URL.Path,
	}
}

//根据key获取解析后的参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

//POST方法的PostForm
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

//GET方法的Query
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

//设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//往响应header中添加kv
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//String类型返回信息
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values)))
}

//JSON类型返回信息
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-type", "application/json")
	c.Status(code)
	//json编码
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//Data类型返回信息
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

//HTML类型返回信息
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
