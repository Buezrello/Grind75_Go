package solutions

import "reflect"

func IsAnagram(s string, t string) bool {
	return reflect.DeepEqual(createMap(s), createMap(t))
}

func createMap(s string) map[rune]int {
	charMap := make(map[rune]int)
	for _, val := range s {
		charMap[val]++
	}
	return charMap
}
