package simple

import "fmt"

func mySqrt(x int) int {
	y := 1

	for y*y <= x {
		y++
	}

	return y - 1
}

func TestMySqrt() {
	fmt.Println(mySqrt(4))
}
