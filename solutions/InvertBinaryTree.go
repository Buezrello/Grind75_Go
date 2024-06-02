package solutions

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := InvertTree(root.Left)
	right := InvertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
