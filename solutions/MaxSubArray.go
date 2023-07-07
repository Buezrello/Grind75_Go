package solutions

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	max := math.MinInt

	for _, num := range nums {
		sum += num
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
