package simple

var r = map[int]int{}

func tribonacci(n int) int {
	if mem, ok := r[n]; ok {
		return mem
	}
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		res := tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
		r[n] = res
		return res
	}
}
