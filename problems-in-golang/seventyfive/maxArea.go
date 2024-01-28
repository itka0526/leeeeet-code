package seventyfive

func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		w := r - l
		h := min(height[l], height[r])
		if w*h > max {
			max = w * h
		}
		if height[l] >= height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}
