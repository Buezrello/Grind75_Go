package solutions

var diameter int

func DiameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	diameterCalculating(root)
	return diameter
}

func diameterCalculating(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := diameterCalculating(node.Left)
	right := diameterCalculating(node.Right)

	if diameter < left+right {
		diameter = left + right
	}

	if left < right {
		return right + 1
	}

	return left + 1
}
