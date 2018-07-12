package src

import (

)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
    if root.Left == nil && root.Right == nil {
	    return 1
    }
    left, right := 0, 0
    if root.Left != nil {
	    left = maxDepth(root.Left)
    }
    if root.Right != nil {
	    right = maxDepth(root.Right)
    }
    if left > right {
	    return left + 1
    } else {
	    return right + 1
    }
}