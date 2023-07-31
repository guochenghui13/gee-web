package gee

import (
	"fmt"
	"strings"
)

type node struct {
	part     string
	children []*node

	pattern string // leave node is not nil
	isWild  bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{part = %s, pattern = %s, isWild = %t}", n.part, n.pattern, n.isWild)
}

func (n *node) matchChild(part string) *node {
	for _, c := range n.children {
		if c.part == part || c.isWild {
			return c
		}
	}
	return nil
}

func (n *node) mathChildren(part string) []*node {
	res := make([]*node, 0)

	for _, child := range n.children {
		if child.isWild || child.part == part {
			res = append(res, child)
		}
	}

	return res
}

// Insert
func (n *node) Insert(pattern string, parts []string, idx int) {
	if idx == len(parts) {
		n.pattern = pattern
		return
	}

	s := parts[idx]
	child := n.matchChild(s)

	if child == nil {
		child = &node{
			part:     s,
			children: make([]*node, 0),
			isWild:   s[0] == '*' || s[0] == ':',
		}
		n.children = append(n.children, child)
	}

	child.Insert(pattern, parts, idx+1)
}

// Search
func (n *node) Search(parts []string, idx int) *node {
	if len(parts) == idx || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	s := parts[idx]
	children := n.mathChildren(s)

	for _, child := range children {
		res := child.Search(parts, idx+1)
		if res != nil {
			return res
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}

	for _, child := range n.children {
		child.travel(list)
	}
}
