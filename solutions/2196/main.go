// create-binary-tree-from-descriptions

package main

import (
	. "leetcode-go/ds"
	"leetcode-go/utils"
)

func main() {
	input := "[[1,2,1],[2,3,0],[3,4,1]]"
	descriptions := utils.JSONParse[[][]int](input)
	root := createBinaryTree(*descriptions)
	PrintBinaryTree(root)
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	isRootMap := make(map[int]bool)

	getNode := func(value int) *TreeNode {
		node, ok := nodeMap[value]
		if !ok {
			node = &TreeNode{Val: value}
			nodeMap[value] = node
		}
		return node
	}

	for _, description := range descriptions {
		parentValue, childValue, isLeft := description[0], description[1], description[2]
		parent := getNode(parentValue)
		child := getNode(childValue)
		if isLeft == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}

		_, ok := isRootMap[parentValue]
		if !ok {
			isRootMap[parentValue] = true
		}
		isRootMap[childValue] = false
	}
	for value, isRoot := range isRootMap {
		if isRoot {
			return getNode(value)
		}
	}
	return nil
}
