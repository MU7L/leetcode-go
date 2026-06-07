package models

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BFSBinaryTreeVisitInfo struct {
	Depth  int
	Parent *TreeNode
	IsLeft bool
}

func PrintBinaryTree(root *TreeNode) {
	lastDepth := 0
	BFSBinaryTree(root, func(node *TreeNode, info *BFSBinaryTreeVisitInfo) {
		if lastDepth != info.Depth {
			lastDepth = info.Depth
			fmt.Print("\n")
		}
		if info.Parent == nil {
			fmt.Printf("%d ", node.Val)
		} else if info.IsLeft {
			fmt.Printf("%d(%dL) ", node.Val, info.Parent.Val)
		} else {
			fmt.Printf("%d(%dR) ", node.Val, info.Parent.Val)
		}
	})
	fmt.Println()
}

func BFSBinaryTree(root *TreeNode, visit func(node *TreeNode, info *BFSBinaryTreeVisitInfo)) {
	type BFSQueueItem struct {
		node   *TreeNode
		parent *TreeNode
		isLeft bool
	}
	queue := []*BFSQueueItem{{node: root}}

	for depth := 0; len(queue) > 0; depth++ {
		levelSize := len(queue) // 当前层的节点数

		// 遍历当前层所有节点
		for range levelSize {
			curr := queue[0]
			queue = queue[1:] // 出队
			visit(curr.node, &BFSBinaryTreeVisitInfo{
				Depth:  depth,
				Parent: curr.parent,
				IsLeft: curr.isLeft,
			})
			// 将下一层节点入队
			if curr.node.Left != nil {
				queue = append(queue, &BFSQueueItem{
					node:   curr.node.Left,
					parent: curr.node,
					isLeft: true,
				})
			}
			if curr.node.Right != nil {
				queue = append(queue, &BFSQueueItem{
					node:   curr.node.Right,
					parent: curr.node,
					isLeft: false,
				})
			}
		}
	}
}
