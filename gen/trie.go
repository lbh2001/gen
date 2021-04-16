package gen

import "strings"

/**
 * @Author: lbh
 * @Date: 2021/4/10
 * @Description: Trie is a data structure to register and parse/search routers.
 */

//pattern:	从根结点到该结点的全路径
//part:		该结点的部分路径
//children:	孩子结点(们)
//isWild:	是否模糊匹配(part中含有":"或"*"则代表模糊匹配)
type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// 获取第一个匹配成功的结点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
		if child.isWild || part[0] == '*' || part[0] == ':' {
			panic("路径冲突：已存在通配路径或普通路径！")
		}
	}
	return nil
}

//基于结点n开始插入新结点
func (n *node) insertChild(pattern string, parts []string, height int) {

	//递归结束条件:遍历完最底层结点(遍历完parts)
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]

	child := n.matchChild(part)

	//结点n的孩子结点中没有part,则插入
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	//递归插入
	child.insertChild(pattern, parts, height+1)
}

// 获取孩子结点中 “part” 为 part 的所有结点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//基于结点n开始匹配结点
//用于获取最终 “pattern” 不为空的结点
func (n *node) search(parts []string, height int) *node {

	//递归结束条件
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	//遍历所有 “part” 为 part 的孩子结点
	for _, child := range children {
		//递归查找
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
