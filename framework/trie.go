package framework

import (
	"errors"
	"strings"
)

const (
	// ParamStart the character in string representation where the underline router starts its dynamic named parameter.
	ParamStart = ":"
	// WildcardParamStart the character in string representation where the underline router starts its dynamic wildcard
	// path parameter.
	WildcardParamStart = "*"

	pathSep  = "/"
)


type Tree struct {
	root *node //根节点
}

type node struct {
	isLast bool //该节点是否能成为一个独立的uri,是否自身就是一个终极节点
	segment string //uri的字符串
	handler ControllerHandler //控制器
	childs []*node //子节点
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
}

//判断一个segment是否是通用segment,即以：开头
func isWildSegment(segment string) bool  {
	return strings.HasPrefix(segment, ParamStart)
}

//过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node  {
	if len(n.childs) == 0 {
		return nil
	}

	//如果segment 是通配符，则所有下一层节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	//过滤所有的下一层子节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			//如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, cnode)
		}else if cnode.segment == segment {
			//如果下一层字节点没有通配符，但是文本完全匹配，则满足需求
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

// 判断路由是否已经在节点的所有子节点树种存在
func (n *node) mathNode(uri string) *node  {
	//使用分割符将uri切割为两部分
	segments := strings.SplitN(uri, pathSep, 2)
	// 第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	//匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)
	//如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在，之间返回nil
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	//如果只有一个segment, 则是最后一个标记
	if len(segments) == 1 {
		//如果segment已经是最后一个节点，判断这些cnode是否是isLast标记
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		//都不是最后一个节点
		return nil
	}

	// 如果有2个segment, 递归每个子节点继续进行查找
	for _, tn := range cnodes {
		tnMatch := tn.mathNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}

	return nil
}

//增加路由节点，路由节点先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name/:age(冲突)
 */
func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root
	if n.mathNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}

	segments := strings.Split(uri, pathSep)
	//对每个segment
	for index, segment := range segments {
		//最终进入Node segment的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments) - 1

		var objNode *node //标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		//如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			//创建一个当前的node 节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
    return nil
}

// 匹配uri
func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.mathNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}