package arraysandhashing

func finalValueAfterOperations(operations []string) int {
	i := 0
	for _, o := range operations {
		switch o {
		case "--X":
			i--
		case "X--":
			i--
		case "++X":
			i++
		case "X++":
			i++
		}
	}
	return i
}
