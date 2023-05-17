package solutions

func ClimbStairs(n int) int {
	if n < 2 {
		return n
	}

	prevPrev, prev := 1, 2
	for i := 3; i < n+1; i++ {
		prevPrev, prev = prev, prevPrev+prev
	}

	return prev
}
