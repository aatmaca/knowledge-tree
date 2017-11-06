package main

import (
	"errors"
	"strings"
)

// Mastery is the interface that represents a user's mastery tree.
type Mastery interface {
	// EncodedPath returns all the paths in the Mastery tree in encoded form.
	EncodedPath() []string

	// FindNode returns the MasteryNode at path p. If not found, ErrPathNotFound is returned.
	FindNode(path string) (Node, error) // ????????

	// Nodes returns all Nodes sorted by Score descending order
	Nodes() []Node

	// Update updates the tree path (and its parents) using information in val (right now just pass nil as value)
	Update(path string, val interface{}) error // ????????

	// SortTree sorts the tree by Score if sort is - in descending order,if sort is + in ascending order.
	SortTree(sort string) // ??????????????
}

// MasteryTree interface will be used by
type MasteryTree struct {
	ep          EncodedPath
	root        *NodeStruct
	nodes       []Node
	nodeStructs []*NodeStruct
}

// EncodedPath returns all the paths in the Mastery tree in encoded form.
func (mt MasteryTree) EncodedPath() []string {
	arr, _ := mt.ep.Decode()
	return arr
}

// FindNode returns the MasteryNode at path p. If not found, ErrPathNotFound is returned.
func (mt MasteryTree) FindNode(path string) (Node, error) {

	var pathArray = (strings.Split(path, ":"))[1:]
	var currentNode = mt.root
	for _, v2 := range pathArray {
		var found bool
		var children = currentNode.Children
		for _, v3 := range children {
			if v3.Value == v2 {
				currentNode = v3
				found = true
				break
			}
		}

		if found == false {
			return nil, errors.New("ErrPathNotFound")
		}
	}

	return currentNode, nil
}

// SortTree sorts the tree by Score if sort is - in descending order,if sort is + in ascending order.
func (mt MasteryTree) SortTree(sortParam string) {
	mt.root.SortMyChildrenRecursively(sortParam)
}

// Update updates the tree path (and its parents) using information in val (right now just pass nil as value)
func (mt MasteryTree) Update(path string, val interface{}) error {

	var pathArray = (strings.Split(path, ":"))[1:]
	var currentNode = mt.root
	currentNode.IncrementScore(1)
	for _, v2 := range pathArray {
		var found bool
		var children = currentNode.Children
		for _, v3 := range children {
			if v3.Value == v2 {
				currentNode = v3
				currentNode.IncrementScore(1)
				found = true
				break
			}
		}

		if found == false {
			return errors.New("ErrPathNotFound")
		}
	}

	return nil
}

// Nodes returns all Nodes sorted by Score descending order
func (mt MasteryTree) Nodes() []Node {

	decreasingOrder := func(p1, p2 *NodeStruct) bool {
		return p1.scoreValue > p2.scoreValue
	}
	By(decreasingOrder).Sort(mt.nodeStructs)

	var nodeArray []Node
	for _, v := range mt.nodeStructs {
		nodeArray = append(nodeArray, v)
	}
	return nodeArray
}

// NewMastery Implement a function that initializes the tree
func NewMastery(ep EncodedPath) (Mastery, error) {

	paths, _ := EncodedPath(ep).Decode()
	nodes, nodeStructs, root, _ := pathArrayToNodeArray(paths)
	var m Mastery = MasteryTree{ep, root, nodes, nodeStructs}

	return m, nil
}
