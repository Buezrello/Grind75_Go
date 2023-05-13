package solutions

import "math"

func IsBalanced(root *TreeNode) bool {
	return heightBalance(root) >= 0
}

func heightBalance(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	leftHeight := heightBalance(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := heightBalance(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if math.Abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return math.Max(leftHeight, rightHeight) + 1

}
