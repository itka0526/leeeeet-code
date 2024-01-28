package seventyfive

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	for i := range res {
		res[i] = true
	}
	for i := range candies {
		for j := range candies {
			if candies[i]+extraCandies < candies[j] {
				res[i] = false
			}
		}
	}
	return res
}

func TestKidsWithCandies() {
	candies := []int{2, 3, 5, 1, 3}

	fmt.Println(kidsWithCandies(candies, 3)) // Expected: [true,true,true,false,true]
}
