package solutions

func LengthOfLongestSubstring(s string) int {
	mapChars := map[uint8]int{}
	maxLength := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := mapChars[ch]; ok {
			i = mapChars[ch]
			clear(mapChars)
		} else {
			mapChars[ch] = i
			if len(mapChars) > maxLength {
				maxLength = len(mapChars)
			}
		}
	}
	return maxLength
}
