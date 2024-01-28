package seventyfive

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	// Check for each position in the array
	// If current position is already filled or neighbors are 1 we cannot plant
	// If a plot is available n decrements and we plant
	// By the end of our iteration if n == 0 we return true else false
	for i := range flowerbed {
		if n == 0 {
			break
		}
		na := flowerbed[i] == 1
		adj := i != 0 && flowerbed[i-1] == 1 || i != len(flowerbed)-1 && flowerbed[i+1] == 1
		if na || adj {
			continue
		}
		flowerbed[i] = 1
		n--
	}
	return n == 0
}

func TestCanPlaceFlowers() {
	f := []int{0, 0, 1, 0, 0}
	fmt.Println("Expected: true; Got: ", canPlaceFlowers(f, 1))
}
