package solutions

func LongestPalindrome(s string) int {
	charCountMap := createCountMap(s)
	odd := false
	palindromeLength := 0

	for _, value := range charCountMap {
		mod := value % 2
		if mod == 1 {
			odd = true
		}
		palindromeLength += value - mod
	}

	mod := len(s) % 2
	if !odd && mod == 1 {
		odd = true
	}

	if odd {
		palindromeLength++
	}

	return palindromeLength
}

func createCountMap(s string) map[int32]int {
	charCountMap := make(map[int32]int)
	for _, chr := range s {
		_, ok := charCountMap[chr]
		if ok {
			charCountMap[chr]++
		} else {
			charCountMap[chr] = 1
		}
	}
	return charCountMap
}
