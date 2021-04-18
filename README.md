<div align='center' ><h1>Gen</h1></div>

<p align="center">English | <a href="https://github.com/lbh2001/gen/blob/master/README_CN.md">简体中文</a></p>

<p align="center"><a href="https://github.com/lbh2001/gen">Gen</a> is a web framework analogous to a simplified <a href="https://github.com/gin-gonic/gin">Gin</a></p>





# Installation

1. go get it

`go get -u github.com/lbh2001/gen/gen`

2. import it

`import "net/http"` (we need package `net/http` to support it)

`import "github.com/lbh2001/gen/gen"`





# Features

- It can handle requests in different ways.(GET, POST ...)
- Using router groups to unify managing paths.
- Completed the redirection of trailing slash paths.
- It stores and matches paths and handler functions based on trie.
- Supporting for global middlewares and custom local middlewares.
- Encapsulates multiple return type information.
- Clever error handling mechanism.
- Supporting HTML template rendering.