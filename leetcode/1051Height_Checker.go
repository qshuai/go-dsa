package leetcode

// https://leetcode.com/problems/height-checker/

// heightChecker
func heightChecker(heights []int) int {
	min, max := heights[0], heights[0]
	for _, height := range heights {
		if height < min {
			min = height
		}
		if height > max {
			max = height
		}
	}

	buckets := make([]int, max-min+1)
	for _, height := range heights {
		buckets[height-min]++
	}

	var offset, missed int
	for idx, item := range buckets {
		for item > 0 {
			if heights[offset] != idx+min {
				missed++
			}

			item--
			offset++
		}
	}

	return missed
}
