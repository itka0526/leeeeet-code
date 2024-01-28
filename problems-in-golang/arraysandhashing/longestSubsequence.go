package arraysandhashing

import "fmt"

func LongestSubsequence() {
	// Idea:
	// 	Subsq length starts from 1 and goes all the way through
	// 	Then length goes up by 1
	// 	And compare the previous maximum subsq sum

	a := []int{0, -11, 20, -19, 8, -19, -12, 14} // 22

	max := 0
	n := 1

	for n <= len(a) {
		for i := range a {
			sum := 0
			l := i + n
			for l > len(a) {
				l--
			}
			fmt.Println("Current: ", a[i:l])
			for j := range a[i:l] {
				sum += a[i:l][j]
			}
			if max < sum {
				max = sum
			}
			fmt.Println(sum)
		}
		n++
	}

	fmt.Println("MAX: ", max)
}
