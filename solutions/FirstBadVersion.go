package solutions

func firstBadVersion(n int) int {
	var middle int
	for start := 1; start <= n; {
		middle = start + (n-start)/2
		if isBadVersion(middle) && !isBadVersion(middle-1) {
			return middle
		} else if isBadVersion(middle) {
			n = middle - 1
		} else {
			start = middle + 1
		}
	}

	return 1
}

func isBadVersion(version int) bool {
	return false
}
