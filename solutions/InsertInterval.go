package solutions

func Insert(intervals [][]int, newInterval []int) [][]int {
	var intervalBefore [][]int
	var intervalAfter [][]int
	start := newInterval[0]
	end := newInterval[1]

	for _, interval := range intervals {
		currentStart := interval[0]
		currentEnd := interval[1]
		if currentEnd < start {
			intervalBefore = append(intervalBefore, interval)
		} else if currentStart > end {
			intervalAfter = append(intervalAfter, interval)
		} else {
			if currentStart < start {
				start = currentStart
			}
			if currentEnd > end {
				end = currentEnd
			}
		}
	}

	var result [][]int
	result = append(result, intervalBefore...)
	result = append(result, []int{start, end})
	result = append(result, intervalAfter...)
	return result
}
