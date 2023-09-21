package slidingwindow

import (
	"github.com/itka0526/problems-in-golang/functions"
)

// not a sliding window problem? Two pointer problem? I am so confused right now... Maybe it derives from two pointer method
// therefore, its called sliding window
func maxProfit(prices []int) int {
	max := 0
	left, right := 0, 1

	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
		} else {
			profit := prices[right] - prices[left]
			if max < profit {
				max = profit
			}
			right++
		}
	}

	return max
}

// 201/212 damn... its not the sliding window approach I don't get it. How am I stuck on the simplest task
// func maxProfit(prices []int) int {
// 	max := 0

// 	for i := 0; i < len(prices); i++ {
// 		j := i + 1
// 		for j < len(prices) {
// 			if prices[i] > prices[j] {
// 				j++
// 				continue
// 			}
// 			profit := prices[j] - prices[i]
// 			if max < profit {
// 				max = profit
// 			}
// 			j++
// 		}
// 	}

// 	return max
// }

// better naive approach 203/212... shit, is there another new method to solve these?
// func maxProfit(prices []int) int {
// 	profit := 0
// 	for i := range prices {
// 		max := maxIndex(prices[i:])
// 		if profit < prices[max+i]-prices[i] {
// 			profit = prices[max+i] - prices[i]FF
// 		}
// 	}
// 	return profit
// }

// func maxIndex(s []int) int {
// 	idx := 0
// 	m := s[idx]

// 	for i := range s {
// 		if m < s[i] {
// 			m = s[i]
// 			idx = i
// 		}
// 	}

// 	return idx
// }

// naive approach does not work here...
// func maxProfit(prices []int) int {
// 	min := minIndex(prices)
// 	max := maxIndex(prices[min:])

// 	return prices[max+min] - prices[min]
// }

// func minIndex(s []int) int {
// 	idx := 0
// 	m := s[idx]

// 	for i := range s {6
// 		if m > s[i] {
// 			m = s[i]
// 			idx = i
// 		}
// 	}

// 	return idx
// }

// func maxIndex(s []int) int {
// 	idx := 0
// 	m := s[idx]

// 	for i := range s {
// 		if m < s[i] {
// 			m = s[i]
// 			idx = i
// 		}
// 	}

// 	return idx
// }

func TestMaxProfit() {
	functions.PrintSolution(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	// functions.PrintSolution(maxProfit([]int{1, 2}))
}
