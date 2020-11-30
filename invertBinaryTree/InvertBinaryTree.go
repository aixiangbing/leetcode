package invertBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertBinaryTree(root *TreeNode)  *TreeNode  {
	if root != nil {
		left := invertBinaryTree(root.Left)
		right := invertBinaryTree(root.Right)

		root.Left = right
		root.Right = left
	}
	return root
}

