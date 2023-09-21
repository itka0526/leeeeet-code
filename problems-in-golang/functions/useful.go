package functions

import "fmt"

func StfuCompiler(a ...interface{}) {}

func PrintSolution(a ...interface{}) {
	for b := range a {
		fmt.Print("Solution: ", a[b], "\n\n")
	}
}
