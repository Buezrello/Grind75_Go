package solutions

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for idx, el := range nums {
		val, ok := hashMap[el]
		if ok {
			return []int{val, idx}
		} else {
			hashMap[target-el] = idx
		}
	}

	return nil
}
