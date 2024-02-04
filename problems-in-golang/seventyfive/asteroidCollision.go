package seventyfive

import "fmt"

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	i := 0

o:
	for i < len(asteroids) {
		for len(res) > 0 && res[len(res)-1] > 0 && asteroids[i] < 0 {
			if res[len(res)-1] > -asteroids[i] {
				i++
				continue o
			} else if res[len(res)-1] == -asteroids[i] {
				res = res[:len(res)-1]
				i++
				continue o
			}
			res = res[:len(res)-1]
		}
		res = append(res, asteroids[i])
		i++
	}
	return res
}

func TestAsteroidCollision() {
	fmt.Println("Exp: [-2,-1, 1, 2]; Got: ", asteroidCollision([]int{-2, -1, 1, 2}))
	fmt.Println("Exp: [10]; Got: ", asteroidCollision([]int{10, 2, -5}))
	fmt.Println("Exp: [-2, -2, -2]; Got: ", asteroidCollision([]int{-2, 1, -2, -2}))
	fmt.Println("Exp: []; Got: ", asteroidCollision([]int{8, -8}))
	fmt.Println("Exp: [5, 10]; Got: ", asteroidCollision([]int{5, 10, -5}))
}
