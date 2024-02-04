package seventyfive

import "fmt"

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, n := range arr {
		m[n]++
	}
	nm := map[int]int{}
	for k, v := range m {
		nm[v] = k
	}
	return len(nm) == len(m)
}

func TestUniqueOccurrences() {
	fmt.Println("Exp: false; Got: ", uniqueOccurrences([]int{1, 2}))
	fmt.Println("Exp: true; Got: ", uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
	fmt.Println("Exp: true; Got: ", uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}
