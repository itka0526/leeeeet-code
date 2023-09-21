package stack

import "github.com/itka0526/problems-in-golang/functions"

func dailyTemperaturesSolution(temperatures []int) []int {
	answers := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		j := i + 1

		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			if answers[j] <= 0 {
				break
			}
			j += answers[j]
		}

		if j < len(temperatures) && temperatures[j] > temperatures[i] {
			answers[i] = j - i
		}
	}

	return answers
}

//	0  1  2  3  4  5  6  7
//
// [0, 0, 0, 0, 0, 1, 0, 0]
func TestDailyTemperaturesSolution() {
	functions.PrintSolution(dailyTemperaturesSolution([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
