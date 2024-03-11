package codeforces

import "fmt"

func PrefixSum() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	i := 0
	for i < n {
		fmt.Scanf("%d", &a[i])
		i++
	}
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		b[i+1] = a[i] + b[i]
	}
	for _, e := range b {
		fmt.Print(e, " ")
	}
	fmt.Print("\n")
}
