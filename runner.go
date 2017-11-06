package main

import (
	"fmt"
	"strings"
)

func main() {

	var sample = "a(b(c,d),e)"
	fmt.Println(sample)

	paths, _ := EncodedPath(sample).Decode()
	fmt.Println(paths)

	ep, _ := Encode(paths)
	fmt.Println(ep)

	fmt.Println("\n---------------Test mastery------------------")
	m, _ := NewMastery(ep)
	fmt.Println(m.EncodedPath())
	fmt.Println(m.Nodes())
	m.SortTree("+")
	fmt.Println(m)
	var node, _ = m.FindNode("a:kkk:tt")
	fmt.Println(node)
	m.Update("a:kkk:tt", nil)
	fmt.Println(m.Nodes())
}

//PrefixSubtree will be used for Decode operation.
type PrefixSubtree struct {
	prefix  string
	subtree string
}

// EncodedPath stores the paths in a tree in a string in encoded form.
type EncodedPath string

// Decode decodes the EncodedPath and returns a slice of leaf paths.
func (ep EncodedPath) Decode() ([]string, error) {

	var result []string
	var arr []PrefixSubtree
	var encodedPath = string(ep)

	var index = strings.Index(encodedPath, "(")
	var root = encodedPath[:index]
	result = append(result, root)
	encodedPath = encodedPath[index+1:len(encodedPath)-1] + ","

	var splittedPath = strings.Split(encodedPath, "")
	var depth int
	var nextSiblingStartIndex int
	for i, v := range splittedPath {
		if v == "(" {
			depth++
		} else if v == "," && depth == 0 {
			var subtree = encodedPath[nextSiblingStartIndex:i]
			arr = append(arr, PrefixSubtree{root, subtree})
			nextSiblingStartIndex = i + 1
		} else if v == ")" {
			depth--
		}
	}

	_, b := DecodeRecursion(arr, result)

	return b, nil
}

// DecodeRecursion is used by Decode method
func DecodeRecursion(currentArr []PrefixSubtree, result []string) ([]PrefixSubtree, []string) {

	var arr []PrefixSubtree

	for _, v := range currentArr {
		var encodedSubPath = v.subtree

		var index = strings.Index(encodedSubPath, "(")

		// This node can be a leaf node.
		if index == -1 {
			result = append(result, v.prefix+":"+encodedSubPath)
			continue
		}

		var subRoot = encodedSubPath[:index]
		var prefix = v.prefix + ":" + subRoot
		result = append(result, prefix)

		encodedSubPath = encodedSubPath[index+1:len(encodedSubPath)-1] + ","

		var splittedPath = strings.Split(encodedSubPath, "")
		var depth int
		var nextSiblingStartIndex int
		for i, v := range splittedPath {
			if v == "(" {
				depth++
			} else if v == "," && depth == 0 {
				var subtree = encodedSubPath[nextSiblingStartIndex:i]
				arr = append(arr, PrefixSubtree{prefix, subtree})
				nextSiblingStartIndex = i + 1
			} else if v == ")" {
				depth--
			}
		}
	}
	if len(arr) == 0 {
		return arr, result
	}

	return DecodeRecursion(arr, result)
}

// Encode is a function that takes a slice of paths and encodes them into an EncodedPath.
func Encode(paths []string) (EncodedPath, error) {

	_, _, root, _ := pathArrayToNodeArray(paths)

	return EncodedPath(AppendChildren(root)), nil
}

// Encode is a function that takes a slice of paths and encodes them into an EncodedPath.
func pathArrayToNodeArray(paths []string) ([]Node, []*NodeStruct, *NodeStruct, error) {

	var nodes []Node
	var nodeStructs []*NodeStruct

	var index = strings.Index(paths[0], ":")
	var rootValue string
	if index == -1 {
		rootValue = paths[0]
	} else {
		rootValue = paths[0][:index]
	}

	var root = &NodeStruct{rootValue, 0, nil, []*NodeStruct{}}
	nodes = append(nodes, root)
	nodeStructs = append(nodeStructs, root)

	for _, v1 := range paths {
		var pathArray = (strings.Split(v1, ":"))[1:]

		var currentNode = root
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
				var child = &NodeStruct{v2, 0, currentNode, []*NodeStruct{}}
				currentNode.Children = append(currentNode.Children, child)
				nodes = append(nodes, child)
				nodeStructs = append(nodeStructs, child)
				//fmt.Println(*currentNode, *child)
			}
		}
	}

	return nodes, nodeStructs, root, nil
}

// AppendChildren recursively constructs EncodedPath
func AppendChildren(child *NodeStruct) string {

	var appendedStr = child.Value

	if len(child.Children) == 0 {
		return appendedStr
	}

	appendedStr += "("
	for _, v := range child.Children {
		appendedStr += AppendChildren(v) + ","
	}
	return appendedStr[:len(appendedStr)-1] + ")"
}
