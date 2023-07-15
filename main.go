package main

import (
	"fmt"
	"grind75/solutions"
)

func main() {
	//fmt.Println(solutions.TwoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(solutions.TwoSum([]int{3, 2, 4}, 6))
	//fmt.Println(solutions.TwoSum([]int{3, 3}, 6))
	//
	//fmt.Println(solutions.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(solutions.MaxProfit([]int{7, 6, 4, 3, 1}))
	//
	//fmt.Println(solutions.IsValid("()"))
	//fmt.Println(solutions.IsValid("()[]{}"))
	//fmt.Println(solutions.IsValid("(]"))
	//fmt.Println(solutions.IsValid("(){}}{"))
	//fmt.Println(solutions.IsValid("(])"))

	//fmt.Println(solutions.IsPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(solutions.IsPalindrome("race a car"))
	//fmt.Println(solutions.IsPalindrome(" "))
	//fmt.Println(solutions.IsPalindrome(".,"))

	//fmt.Println(solutions.IsAnagram("anagram", "nagaram"))
	//fmt.Println(solutions.IsAnagram("rat", "cat"))

	//fmt.Println(solutions.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	//fmt.Println(solutions.Search([]int{-1, 0, 3, 5, 9, 12}, 2))

	//image1 := [][]int{
	//	{1, 1, 1},
	//	{1, 1, 0},
	//	{1, 0, 1},
	//}
	//
	//image2 := [][]int{
	//	{0, 0, 0},
	//	{0, 0, 0},
	//}
	//
	//fmt.Println(solutions.FloodFill(image1, 1, 1, 2))
	//fmt.Println(solutions.FloodFill(image2, 0, 0, 0))

	//fmt.Println(solutions.CanConstruct("aac", "acab"))

	//fmt.Println(solutions.ClimbStairs(5))

	//fmt.Println(solutions.LongestPalindrome("abccccdd"))
	//fmt.Println(solutions.LongestPalindrome("a"))
	//fmt.Println(solutions.LongestPalindrome("abcdadcb"))

	//fmt.Println(solutions.MajorityElement([]int{3, 2, 3}))
	//fmt.Println(solutions.MajorityElement([]int{2, 2, 1, 1, 2, 2, 1}))

	//fmt.Println(solutions.AddBinary("11", "1"))
	//fmt.Println(solutions.AddBinary("1010", "1011"))
	//fmt.Println(solutions.AddBinary("0", "0"))

	//root := &solutions.TreeNode{Val: 1, Left: nil, Right: nil}
	//root.Left = &solutions.TreeNode{Val: 2, Left: nil, Right: nil}
	//root.Left.Left = &solutions.TreeNode{Val: 4, Left: nil, Right: nil}
	//root.Left.Right = &solutions.TreeNode{Val: 5, Left: nil, Right: nil}
	//root.Right = &solutions.TreeNode{Val: 3, Left: nil, Right: nil}
	//
	//fmt.Println(solutions.DiameterOfBinaryTree(root))
	//
	//root = &solutions.TreeNode{Val: 1, Left: nil, Right: nil}
	//root.Left = &solutions.TreeNode{Val: 2, Left: nil, Right: nil}
	//
	//fmt.Println(solutions.DiameterOfBinaryTree(root))

	//fmt.Println(solutions.ContainsDuplicate([]int{1, 2, 3, 1}))
	//fmt.Println(solutions.ContainsDuplicate([]int{1, 2, 3, 4}))
	//fmt.Println(solutions.ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	fmt.Println(solutions.Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(solutions.Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
