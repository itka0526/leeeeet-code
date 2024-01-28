package seventyfive

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	s := []byte{}
	for i := 0; i < len(chars); i++ {
		j := i + 1
		k := 1
		for j < len(chars) && chars[i] == chars[j] {
			j++
			k++
		}
		if k > 1 {
			b := []byte(strconv.Itoa(k))
			s = append(s, chars[i])
			s = append(s, b...)
		} else {
			s = append(s, chars[i])
		}
		i = j - 1
	}
	copy(chars, s)
	return len(s)
}

func TestCompress() {
	i := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println("Exp: '6'; Got: ", compress(i))
	fmt.Println("Should be: \t a2b2c3 ")
	fmt.Print("Got: \t\t ")
	for _, e := range i {
		fmt.Print(string(e))
	}
}
