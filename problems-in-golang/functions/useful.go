package functions

import "fmt"

func StfuCompiler(a ...interface{}) {}

func PrintSolution(a ...interface{}) {
	for b := range a {
		fmt.Print("Solution: ", a[b], "\n\n")
	}
}

func LogAnswer(
	got any,
	exp any,
) {
	if _, err := fmt.Println("Answer: ", got, "; Exp: ", exp); err != nil {
		panic("Cannot log this!")
	}
}
