package main

import (
	"fmt"
)

// Node interface will be used by
type Node interface {
	// Score returns how many times this node has been updated
	Score() float64

	// IncrementScore increments score by given value
	IncrementScore(a float64)

	// SortMyChildrenRecursively calls SortMyChildren recursively
	SortMyChildrenRecursively(sortParam string)

	// SortMyChildren sorts Children by Score if sort is - in descending order, if sort is + in ascending order.
	SortMyChildren(sortParam string)

	// Path returns the path of this node such as "a:b:c" where "c" would be the name of the node and "b" would be parent
	Path() string

	// String is our toString method.
	String() string
}

// NodeStruct will be used by
type NodeStruct struct {
	Value      string
	scoreValue float64 // how many times this node has been updated
	Parent     *NodeStruct
	Children   []*NodeStruct
	//Children list.List
}

// Score returns how many times this node has been updated
func (node *NodeStruct) Score() float64 {
	return node.scoreValue
}

// IncrementScore increments score by given value
func (node *NodeStruct) IncrementScore(a float64) {
	node.scoreValue = node.scoreValue + a
}

// SortMyChildrenRecursively calls SortMyChildren recursively
func (node *NodeStruct) SortMyChildrenRecursively(sortParam string) {
	node.SortMyChildren(sortParam)
	for _, v := range node.Children {
		v.SortMyChildrenRecursively(sortParam)
	}
}

// SortMyChildren sorts Children by Score if sort is - in descending order, if sort is + in ascending order.
func (node *NodeStruct) SortMyChildren(sortParam string) {
	//fmt.Println(node.Children)

	if sortParam == "+" {
		increasingOrder := func(p1, p2 *NodeStruct) bool {
			return p1.scoreValue < p2.scoreValue
		}
		By(increasingOrder).Sort(node.Children)
		//fmt.Println("By increasing order:", node.Children)
	} else {
		decreasingOrder := func(p1, p2 *NodeStruct) bool {
			return p1.scoreValue > p2.scoreValue
		}
		By(decreasingOrder).Sort(node.Children)
		//fmt.Println("By decreasing order:", node.Children)
	}
}

// Path returns the path of this node such as "a:b:c" where "c" would be the name of the node and "b" would be parent
func (node *NodeStruct) Path() string {

	parent := node.Parent

	if parent == nil {
		return node.Value
	}

	return parent.Path() + ":" + node.Value
}

func (node *NodeStruct) String() string {
	return fmt.Sprintf("%s", node.Value)
}
