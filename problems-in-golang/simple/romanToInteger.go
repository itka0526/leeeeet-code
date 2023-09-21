package simple

var d map[uint8]int = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var val, lastVal, currVal int
	for i := len(s) - 1; i >= 0; i-- {
		currVal = d[s[i]]
		if currVal < lastVal {
			val -= currVal
		} else {
			val += currVal
		}
		lastVal = currVal
	}
	return val
}
