package solutions

func Search(nums []int, target int) int {
	return findIndex(nums, target, 0, len(nums)-1)
}

func findIndex(nums []int, target int, start int, end int) int {
	if end < start {
		return -1
	}

	middle := start + ((end - start) / 2)

	if nums[middle] == target {
		return middle
	}

	if target < nums[middle] {
		return findIndex(nums, target, start, middle-1)
	} else {
		return findIndex(nums, target, middle+1, end)
	}
}
