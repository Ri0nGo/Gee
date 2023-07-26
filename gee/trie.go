package gee

import (
	"strings"
)

/*
Pattern 只有是最后一层叶子节点时，才会设置pattern，非叶子节点时不会设置pattern的
isWild 是否精确匹配（确切的说，isWild 用来表示模糊的自定义正则匹配标识），part 含有 : 或 * 时为true
children 标识该节点所包含的子节点
*/
type node struct {
	isWild   bool
	pattern  string // url 全路径字符串，ep: /api/v1/user/login
	part     string // url 路径的一部分，ep: v1 等。
	children []*node
}

// matchChild 查找一个节点，即 传入的part 和 节点的part 相同，且isWild 为true
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// child.isWild 表示为当前节点为动态匹配，类似于:lang, 表示匹配任意字符，并将该字符赋值给lang
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 查询所有匹配的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 将url 路径拆解后依次插入到节点中（多叉树）
/*
pattern url 全路径
parts 将url 以 / 切割之后的字符数组，ep: [api, v1, user, login]
height 标识parts的索引下标，即依次遍历parts中的每个值
*/
func (n *node) insert(pattern string, parts []string, height int) {
	// 遍历到最后一个part部分时，也表示当前节点为叶子节点了。
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 注意，这里千万不能写出 := ，若写成:= 则表示在 if 作用域内定义了一个child，此时外部的child是nil
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search 在parts中从后往前搜索，直到一个node不等于nil的值或返回nil
func (n *node) search(parts []string, height int) *node {
	// 当遍历到最末尾的part时，若part是叶子节点，则返回该node，此时node就为result
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		// 递归遍历，一直到遍历最后的一个part值
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
