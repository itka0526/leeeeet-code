package bitmanipulation

func hammingWeight(n int) (ans int) {
	// 11 1
	// 5  1
	// 2  0
	// 1  1
	// 0  0
	// 0 1 0 1 1
	for n > 0 {
		if n%2 == 1 {
			ans += 1
		}
		n /= 2
	}
	return ans
}
