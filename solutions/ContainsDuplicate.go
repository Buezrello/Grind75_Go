package solutions

func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		} else {
			set[num] = struct{}{}
		}
	}
	return false
}
