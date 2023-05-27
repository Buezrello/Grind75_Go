package solutions

func MajorityElement(nums []int) int {
	count := 0
	var res int

	for _, num := range nums {
		if count == 0 {
			res = num
		}
		count += equalOrNot(res, num)
	}

	return res
}

func equalOrNot(a, b int) int {
	if a == b {
		return 1
	} else {
		return -1
	}
}
