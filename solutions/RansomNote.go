package solutions

import "strings"

func CanConstruct(ransomNote string, magazine string) bool {
	for _, c := range magazine {
		ransomNote = strings.Replace(ransomNote, string(c), "", 1)
	}
	return len(ransomNote) == 0
}
