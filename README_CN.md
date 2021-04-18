<div align='center' ><font size='70'>Gen</font></div>

<p align="center"><a href="/lbh2001/gen/blob/master/README.md">English</a> | 简体中文</p>

<p align="center"><a href="https://github.com/lbh2001/gen">Gen</a>是一个类似于<a href="https://github.com/gin-gonic/gin">Gin</a>简化版的web框架。 </p>





# 下载

1. 下载到本地

`go get -u github.com/lbh2001/gen/gen`

2. 引入到项目

`import "net/http"` (需要 `net/http` 包的支持)

`import "github.com/lbh2001/gen/gen"`





# 特性

- 可以处理不同方式的请求。（GET、POST等）
- 使用路由组统一管理路径。
- 对尾随斜杠路径进行重定向。
- 基于前缀树储存、匹配路径及其处理函数。
- 支持全局与自定义局部中间件。
- 封装了多种返回类型的信息。
- 巧妙的错误处理机制。
- 支持HTML模板渲染功能。