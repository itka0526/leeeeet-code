package seventyfive

import "fmt"

func largestAltitude(gain []int) int {
	ha := 0
	a := 0
	for _, e := range gain {
		a += e
		if a > ha {
			ha = a
		}
	}
	return ha
}

func TestLargestAltitude() {
	fmt.Println("Exp: 781; Got: ", largestAltitude([]int{44, 32, -9, 52, 23, -50, 50, 33, -84, 47, -14, 84, 36, -62, 37, 81, -36, -85, -39, 67, -63, 64, -47, 95, 91, -40, 65, 67, 92, -28, 97, 100, 81}))
	fmt.Println("Exp: 1; Got: ", largestAltitude([]int{-5, 1, 5, 0, -7}))
}
