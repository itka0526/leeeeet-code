package simple

func sumWealth(b []int) int {
	s := 0
	for i := range b {
		s += b[i]
	}
	return s
}

func maximumWealth(accounts [][]int) int {
	wealthiest := sumWealth(accounts[0])
	for i := 1; i < len(accounts); i++ {
		currWealth := sumWealth(accounts[i])
		if wealthiest < currWealth {
			wealthiest = currWealth
		}
	}
	return wealthiest
}
