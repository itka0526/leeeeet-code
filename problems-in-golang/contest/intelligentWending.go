package contest

import (
	"fmt"
)

func intelligentWending() {
	var b, c, r, d int

	// b = 10
	// c = 700000
	// r = 350000
	// d = 200000
	// b = 21
	// c = 1000000
	// r = 1100000
	// d = 0

	if c+d >= 999_999 {
		// has enough to buy another
		fmt.Println((b*1_000_000 + c) / r)
	}

	// z := b*1_000_000 + c
	// a := 0
	// for c >= r && z >= r {
	// 	z -= r
	// 	c -= r
	// 	d += r
	// 	a += 1
	// }

	// for z >= r && z%r == 0 {
	// 	fmt.Println(z, r)
	// 	// it can proceed if only z = r
	// 	z -= r
	// 	a += 1
	// }
	// give c back by subtracting from d
	// fmt.Println(a, d)

	// _, err := fmt.Scanf("%d %d \n %d %d", &b, &c, &r, &d)
	// if err != nil {
	// 	panic("Unexpected... ")
	// }

	// I understand we have to use recursion it seems
	// a := 0
	// var calc func(v int, vv int, vvv int, vvvv int)
	// calc = func(i, j, k, l int) {
	// 	z := i*1_000_000 + j
	// 	if z < r {
	// 		return
	// 	}
	// 	a++
	// 	z -= k
	// 	if z > l {
	// 		return
	// 	}
	// 	l -= z
	// 	calc(z, l, k, l+j)
	// }
	// calc(b, c, r, d)
	// fmt.Println(a)
	// b = b * 10^6
	// c = c * 1
	// r = some amount
	// d = - d * 1
	// amount := 0

	// z := b*1_000_000 + c*1
	// for z >= r {
	// 	z -= r
	// 	amount++
	// }

	// if z >= d {
	// 	d -= z
	// }
	// // for {
	// 	z := b*1_000_000 + c*1
	// 	if z < r {
	// 		// More money
	// 		break
	// 	}
	// 	if z >= r {
	// 		change := z - r
	// 		// Can the wending machine give back change?
	// 		// Give Kvas-Klass
	// 		amount++
	// 		// Give change back
	// 		d -= change
	// 		// Use the 1 rubles as a change for the future purchases
	// 		d += c
	// 	}
	// }

	// fmt.Println(amount, " change left: ", d)
}

func TestIntelligentWending() {
	intelligentWending()
}
