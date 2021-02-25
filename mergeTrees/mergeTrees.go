package mergeTrees

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	merged := &TreeNode{root1.Val + root2.Val, mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)}
	return merged
}



func diameterOfBinaryTree(root *TreeNode) int {

	ans := 1

	depth(root, &ans)

	return ans
}

func depth(root *TreeNode, ans *int) int {

	if root == nil {
		return 0
	}

	l := depth(root.Left, ans)
	r := depth(root.Right, ans)
	*ans = max(*ans, l + r + 1)
	return max(l, r) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}