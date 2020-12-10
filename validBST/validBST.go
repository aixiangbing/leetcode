package validBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	flag := true
	return validate(root, pre, flag)
}

func validate(root *TreeNode, prev int, flag bool)  bool {

	if root == nil {
		return true
	}

	if flag && root.Left != nil {
		validate(root.Left, prev, flag)
	}

	if root.Val <= prev {
		flag = false
	}
	prev = root.Val

	if flag && root.Right != nil {
		validate(root.Right, prev, flag)
	}

	return flag
}


