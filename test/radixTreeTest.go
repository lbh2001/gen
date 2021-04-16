package test

/**
 * @Author: lbh
 * @Date: 2021/4/12
 * @Description: just a try and test.
 */

//结点的类型
const (
	simple   uint = iota //普通结点，默认类型
	root                 //根结点
	param                //带参数的结点 ":"
	catchAll             //匹配所有的结点"*"
)

/*temp start*/
type Context struct{}

/*temp end*/

type HandlerFunc func(c *Context)

type node struct {
	fullPath string      //截止到该结点的全路径
	path     string      //该结点储存的部分路径
	children []*node     //孩子结点们
	nodeType uint        //结点的类型
	indices  []byte      //孩子结点们的首字母
	handlers HandlerFunc //处理函数
}

func (n *node) addRoute(path string, handler HandlerFunc) {
	fullPath := path

	//如果是空树
	if len(n.path) == 0 && len(n.children) == 0 {
		emptyTreeInsert(n, path, fullPath, handler)
		n.nodeType = root
		return
	}

	notEmptyTreeInsert(n, path, fullPath, handler)

}

//当空树时直接插入
func emptyTreeInsert(n *node, path string, fullPath string, handler HandlerFunc) {
	n.insertChild(path, fullPath, handler)
}

//封装gin源码中的
//walk:
//	for{}
//部分
func notEmptyTreeInsert(n *node, path string, fullPath string, handler HandlerFunc) {
	i := longestCommonPrefix(n.path, path)

	//如果 i < len(n.path)
	//则说明n需要分裂
	//因为n.path即当前路径长度大于公共前缀长度
	//如 n.path = /namespace
	//而   path = /name 	(只有n.path要分裂)
	//
	//或者
	//   n.path = /application
	//     path = /agent      (都要分裂)
	//
	if i < len(n.path) {
		nodePathLonger(i, n, path, fullPath)
	}

	//如果 i < len(path)
	//如 n.path = /name
	//而   path = /namespace (只有path要分裂)
	//
	//或者
	//   n.path = /application
	//     path = /agent      (都要分裂)
	//
	if i < len(path) {
		newPathLonger(i, n, path, fullPath)
	}

}

func nodePathLonger(i int, n *node, path string, fullPath string) {
	//child结点继承了n的几乎所有属性
	//child将位于n和n的孩子结点们的中间
	child := &node{
		fullPath: n.fullPath,
		path:     n.path[i:], //非公共的部分
		children: n.children,
		nodeType: n.nodeType,
		indices:  n.indices,
		handlers: n.handlers,
	}

	//n的孩子结点变成child结点
	n.children = []*node{child}
	n.indices = append(n.indices, n.path[i])

	//for example:
	//
	//	 n.path :	/namespace
	// ->n.path :	space
	//n.fullPath:	/usr/namespace
	//可计算验证
	n.path = path[:i]
	n.fullPath = n.fullPath[:len(fullPath)-len(n.path)]
	n.handlers = nil
}

func newPathLonger(i int, n *node, path string, fullPath string) {}

func min(lengthA, lengthB int) int {
	if lengthA > lengthB {
		return lengthB
	} else {
		return lengthA
	}
}

//获得两个字符串的最长公共前缀长度
func longestCommonPrefix(pathA, pathB string) int {
	i := 0
	max := min(len(pathA), len(pathB))
	for i < max && pathA[i] == pathB[i] {
		i++
	}
	return i
}

//插入结点
func (n *node) insertChild(path string, fullPath string, handler HandlerFunc) {}
