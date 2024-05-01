package bitmanipulation

func countBits(n int) (ans []int) {
	ans = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		cnt := 0
		j := i
		for j > 0 {
			if j&1 == 1 {
				cnt++
			}
			j /= 2
		}
		ans[i] = cnt
	}
	return ans
}
